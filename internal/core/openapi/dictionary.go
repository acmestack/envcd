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
	"time"

	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/constants"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/pkg/entity/result"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-gonic/gin"
)

type dictParams struct {
	DictKey   string `json:"dictKey"`
	DictValue string `json:"dictValue"`
	Version   string `json:"version"`
	State     bool   `json:"state"`
}

func (openapi *Openapi) dictionary(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		// get user id from gin context
		userId := stringsx.ToInt(ginCtx.Param("userId"))
		scopeSpaceId := stringsx.ToInt(ginCtx.Param("scopeSpaceId"))
		dictId := stringsx.ToInt(ginCtx.Param("dictId"))
		dict := entity.Dictionary{Id: dictId, UserId: userId, ScopeSpaceId: scopeSpaceId}
		dictionary, err := dao.New(openapi.storage).SelectDictionary(dict)
		if err != nil {
			return result.InternalServerErrorFailure(err.Error())
		}
		return result.Success(dictionary)
	})
}

func (openapi *Openapi) createDictionary(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		param := dictParams{}
		if err := ginCtx.ShouldBindJSON(&param); err != nil {
			fmt.Printf("Bind error, %v\n", err)
			return result.InternalServerErrorFailure(constants.IllegalJsonBinding)
		}
		// get userId and appId from gin context
		userId := stringsx.ToInt(ginCtx.Param("userId"))
		scopeSpaceId := stringsx.ToInt(ginCtx.Param("scopeSpaceId"))

		dictionary := entity.Dictionary{
			UserId:       userId,
			ScopeSpaceId: scopeSpaceId,
			DictKey:      param.DictKey,
			DictValue:    param.DictValue,
			State:        param.State,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		// Strategy
		// scopespace + username + dictKey + version
		// insertDictionary, i, err := dao.New(openapi.storage).InsertDictionary(dictionary)
		// openapi.exchange.Put(dictionary.DictKey, dictionary.DictValue)
		// if err != nil {
		//	return nil
		//}
		fmt.Println(dictionary)
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
		userId := stringsx.ToInt(ginCtx.Param("userId"))
		appId := stringsx.ToInt(ginCtx.Param("appId"))
		dictId := stringsx.ToInt(ginCtx.Param("dictId"))
		dict := entity.Dictionary{Id: dictId, UserId: userId, ScopeSpaceId: appId}
		// query dictionary exist
		daoApi := dao.New(openapi.storage)
		dictionary, err := daoApi.SelectDictionary(dict)
		if err != nil {
			return result.InternalServerErrorFailure(err.Error())
		}
		if len(dictionary) == 0 {
			return result.Failure(constants.DictNotFound, http.StatusBadRequest)
		}
		exchangeErr := openapi.exchange.Remove(getFirstDictionary(dictionary).DictKey)
		if exchangeErr != nil {
			return result.InternalServerErrorFailure(exchangeErr.Error())
		}
		retId, delErr := daoApi.DeleteDictionary(getFirstDictionary(dictionary))
		if delErr != nil {
			return result.InternalServerErrorFailure(delErr.Error())
		}
		return result.Success(retId)
	})
}

func getFirstDictionary(dictionaryList []entity.Dictionary) entity.Dictionary {
	return dictionaryList[0]
}

func (openapi *Openapi) dictionaries(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil
	})
}

func (openapi *Openapi) dictionariesByFuzzyKey(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil
	})
}
