/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"fmt"
	"github.com/acmestack/envcd/internal/core/plugin"
	openservice "github.com/acmestack/envcd/internal/core/service"

	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/acmestack/godkits/gox/errorsx"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(op *openservice.OpenService) func(context2 *gin.Context) {
	return func(context2 *gin.Context) {
		c := &context.Context{Action: func() (*data.EnvcdResult, error) {
			fmt.Println("hello world")
			err := op.Envcd.Put("key", "value")
			if err != nil {
				return nil, err
			}
			return nil, errorsx.Err("test error")
		}}
		if ret, err := plugin.NewChain(op.Executors).Execute(c); err != nil {
			fmt.Printf("ret = %v, error = %v", ret, err)
		}

	}
}
