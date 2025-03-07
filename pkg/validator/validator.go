// Copyright 2020 Talhuang<talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package validator defines iam custom binding validators used by gin.
package validator

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/marmotedu/component-base/pkg/validation"
)

// validateUsername checks if a given username is illegal.
func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	fmt.Println(username)
	if errs := validation.IsQualifiedName(username); len(errs) > 0 {
		return false
	}

	return true
}

// validatePassword checks if a given password is illegal.
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if err := validation.IsValidPassword(password); err != nil {
		return false
	}
	fmt.Println(password)

	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		fmt.Println(ok)
		_ = v.RegisterValidation("username", validateUsername)
		_ = v.RegisterValidation("password", validatePassword)
	}
}
