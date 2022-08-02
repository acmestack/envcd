/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
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

package openapi

import (
	"fmt"
	"net/http"

	"github.com/acmestack/envcd/internal/core/plugin"
	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/acmestack/godkits/gox/errorsx"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-gonic/gin"
)

func (openapi *Openapi) dictionary(ctx *gin.Context) {
	c, _ := buildContext(ctx)
	c.Action = func() (*data.EnvcdResult, error) {
		// get user id from gin context
		userId := stringsx.ToInt(ctx.Param("userId"))
		appId := stringsx.ToInt(ctx.Param("appId"))
		configId := stringsx.ToInt(ctx.Param("configId"))
		dict := entity.Dictionary{Id: configId, UserId: userId, ApplicationId: appId}
		dictionary, err := dao.New(openapi.storage).SelectDictionary(dict)
		if err != nil {
			return data.Failure(err.Error()), err
		}
		return data.Success(dictionary), nil
	}
	ret, err := plugin.NewChain(openapi.executors).Execute(c)
	if err != nil {
		// FIXME log.error("ret = %v, error = %v", ret, err)
		fmt.Printf("ret = %v, error = %v", ret, err)
		ctx.JSON(http.StatusBadRequest, ret.Data)
	}
	ctx.JSON(http.StatusOK, ret.Data)
}

func (openapi *Openapi) putDictionary(ctx *gin.Context) {
	c := &context.Context{Action: func() (*data.EnvcdResult, error) {
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil, errorsx.Err("test error")
	}}
	if ret, err := plugin.NewChain(openapi.executors).Execute(c); err != nil {
		fmt.Printf("ret = %v, error = %v", ret, err)
	}
}

func (openapi *Openapi) removeDictionary(ctx *gin.Context) {
	c := &context.Context{Action: func() (*data.EnvcdResult, error) {
		fmt.Println("hello world")
		// delete config
		// ConfigDao.delete();
		// go LogDao.save()
		// openapi.exchange.Remove("key")
		return nil, errorsx.Err("test error")
	}}
	if ret, err := plugin.NewChain(openapi.executors).Execute(c); err != nil {
		fmt.Printf("ret = %v, error = %v", ret, err)
	}
}
