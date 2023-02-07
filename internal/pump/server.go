// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pump

import (
	"fmt"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis"
	"github.com/marmotedu/log"

	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/apiserver/config"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pump/pumps"
)

var pmps []pumps.Pump

type pumpServer struct {
	secInterval    int
	omitDetails    bool
	mutex          *redsync.Mutex
	analyticsStore storage.AnalyticsStorage
	pumps          map[string]options.PumpConfig
}

// preparedGenericAPIServer is a private wrapper that enforces a call of PrepareRun() before Run can be invoked.
type preparedPumpServer struct {
	*pumpServer
}

func createPumpServer(cfg *config.Config) (*pumpServer, error) {
	// use the same redis database with authorization log history
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisOptions.Host, cfg.RedisOptions.Port),
		Username: cfg.RedisOptions.Username,
		Password: cfg.RedisOptions.Password,
	})

	rs := redsync.New(goredis.NewPool(client))

	server := &pumpServer{
		secInterval:    cfg.PurgeDelay,
		omitDetails:    cfg.OmitDetailedRecording,
		mutex:          rs.NewMutex("iam-pump", redsync.WithExpiry(10*time.Minute)),
		analyticsStore: &redis.RedisClusterStorageManager{},
		pumps:          cfg.Pumps,
	}

	if err := server.analyticsStore.Init(cfg.RedisOptions); err != nil {
		return nil, err
	}

	return server, nil
}

func (s *pumpServer) initialize() {
	pmps = make([]pumps.Pump, len(s.pumps))
	i := 0
	for key, pmp := range s.pumps {
		pumpTypeName := pmp.Type
		if pumpTypeName == "" {
			pumpTypeName = key
		}

		pmpType, err := pumps.GetPumpByName(pumpTypeName)
		if err != nil {
			log.Errorf("Pump load error (skipping): %s", err.Error())
		} else {
			pmpIns := pmpType.New()
			initErr := pmpIns.Init(pmp.Meta)
			if initErr != nil {
				log.Errorf("Pump init error (skipping): %s", initErr.Error())
			} else {
				log.Infof("Init Pump: %s", pmpIns.GetName())
				pmpIns.SetFilters(pmp.Filters)
				pmpIns.SetTimeout(pmp.Timeout)
				pmpIns.SetOmitDetailedRecording(pmp.OmitDetailedRecording)
				pmps[i] = pmpIns
			}
		}
		i++
	}
}
