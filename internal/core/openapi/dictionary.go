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

	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/pkg/entity/result"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-gonic/gin"
)

func (openapi *Openapi) dictionary(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		// get user id from gin context
		userId := stringsx.ToInt(ginCtx.Param("userId"))
		appId := stringsx.ToInt(ginCtx.Param("appId"))
		configId := stringsx.ToInt(ginCtx.Param("configId"))
		dict := entity.Dictionary{Id: configId, UserId: userId, ApplicationId: appId}
		dictionary, err := dao.New(openapi.storage).SelectDictionary(dict)
		if err != nil {
			return result.InternalServerErrorFailure(err.Error())
		}
		return result.Success(dictionary)
	})
}

func (openapi *Openapi) putDictionary(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil
	})
}

func (openapi *Openapi) updateDictionary(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil
	})
}

func (openapi *Openapi) removeDictionary(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// delete config
		// ConfigDao.delete();
		// go LogDao.save()
		// openapi.exchange.Remove("key")
		return nil
	})
}
