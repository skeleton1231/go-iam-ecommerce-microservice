// Copyright 2020 Talhuang<talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/middleware"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/log"
)

// List list all the secrets.
func (s *SecretController) List(c *gin.Context) {
	log.L(c).Info("list secret function called.")
	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	secrets, err := s.srv.Secrets().List(c, c.GetString(middleware.UsernameKey), r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, secrets)
}
