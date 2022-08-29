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

	"github.com/acmestack/envcd/internal/core/exchanger"
	"github.com/acmestack/envcd/internal/core/storage"
	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/internal/pkg/result"
	"github.com/acmestack/godkits/gox/errorsx"
	"github.com/gin-gonic/gin"
)

type PageListVO struct {
	Page      int64       `json:"page"`
	PageSize  int64       `json:"pageSize"`
	Total     int64       `json:"total"`
	TotalPage int64       `json:"totalPage"`
	List      interface{} `json:"list"`
}

type Openapi struct {
	exchange *exchanger.Exchange
	storage  *storage.Storage
	contexts map[string]*context.Context
}

func Start(serverSetting *config.Server, exchange *exchanger.Exchange, storage *storage.Storage) {
	openapi := &Openapi{
		exchange: exchange,
		storage:  storage,
		contexts: map[string]*context.Context{},
	}
	openapi.initServer(serverSetting)
}

// initServer start gin http server
//  @receiver openapi open api
//  @param serverSetting server config
func (openapi *Openapi) initServer(serverSetting *config.Server) {
	gin.SetMode(serverSetting.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverSetting.Port),
		Handler:        openapi.buildRouter(),
		ReadTimeout:    time.Duration(serverSetting.ReadTimeout) * time.Millisecond,
		WriteTimeout:   time.Duration(serverSetting.WriteTimeout) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	// todo log
	//log.Info("start http server listening %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		// todo log
		//log.Error("start http server error %v", err)
	}
}

// todo build Router
func (openapi *Openapi) buildRouter() *gin.Engine {
	router := gin.Default()
	// build context for peer request
	router.Use(openapi.buildContext)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(ginCtx *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			failure := result.InternalFailure(errorsx.Err(fmt.Sprintf("error: %s", err)))
			ginCtx.JSON(failure.HttpStatusCode, failure.Data)
		}
		ginCtx.AbortWithStatus(http.StatusInternalServerError)
	}))

	// login and logout
	router.POST("/login", openapi.login)
	router.GET("/logout", openapi.logout)

	// version 1 group
	v1 := router.Group("/v1")

	// user group routers
	usersGroup := v1.Group("/users")
	{
		// fuzzy filter => ?page=2&name=
		usersGroup.GET("", openapi.users)
		usersGroup.POST("", openapi.createUser)
		usersGroup.PUT("/:userId", openapi.updateUser)
		usersGroup.GET("/:userId", openapi.user)
		usersGroup.DELETE("/:userId", openapi.removeUser)

		// user's all scopeSpaces
		// fuzzy filter => ?page=2&scopespace-name=
		usersGroup.GET("/:userId/scopeSpaces", openapi.userScopeSpaces)

		// user's all dictionaries under one scopespace
		// fuzzy filter => ?page=2&scopespace-name=abc&dictionary-key=aaa
		usersGroup.GET("/:userId/scopespace/:scopeSpaceId/dictionaries", openapi.userDictionariesUnderScopeSpace)

		// user's all dictionaries
		// fuzzy filter => ?page=2&dictionary-key=aaa
		usersGroup.GET("/:userId/dictionaries", openapi.userDictionaries)
	}

	// scopeSpaces group routers
	scopeSpacesGroup := v1.Group("/scopeSpaces")
	{
		// fuzzy filter => ?page=2&user=abc&scopespace-name=
		scopeSpacesGroup.GET("", openapi.scopeSpaces)
		scopeSpacesGroup.POST("", openapi.createScopeSpace)
		scopeSpacesGroup.GET("/:scopeSpaceId", openapi.scopeSpace)
		scopeSpacesGroup.PUT("/:scopeSpaceId", openapi.updateScopeSpace)
		scopeSpacesGroup.DELETE("/:scopeSpaceId", openapi.removeScopeSpace)
	}

	// dictionaries group routers
	dictionariesGroup := v1.Group("/dictionaries")
	{
		// fuzzy filter => ?page=2&user=abc&dictionary-key=
		dictionariesGroup.GET("", openapi.dictionaries)
		dictionariesGroup.POST("", openapi.createDictionary)
		dictionariesGroup.GET("/:dictionaryId", openapi.dictionary)
		dictionariesGroup.PUT("/:dictionaryId", openapi.updateDictionary)
		dictionariesGroup.DELETE("/:dictionaryId", openapi.removeDictionary)
	}

	return router
}

// execute to caller
func (openapi *Openapi) execute(ginCtx *gin.Context, permissionAction context.EnvcdActionFunc, logicAction context.EnvcdActionFunc) {
	requestId := ginCtx.Request.Header.Get(requestIdHeader)
	c := openapi.contexts[requestId]
	ret := result.InternalFailure0()
	if c != nil && c.RequestId == requestId {
		for _, action := range []context.EnvcdActionFunc{openapi.validate(c), permissionAction, logicAction} {
			if action == nil {
				continue
			}
			if ret = action(); ret != nil {
				break
			}
		}
	}
	ginCtx.JSON(ret.HttpStatusCode, ret.Data)
	delete(openapi.contexts, requestId)
}

func (openapi *Openapi) doOperationLogging(userId int, message string) {
	go func() {
		_, _, err := dao.New(openapi.storage).InsertLogging(entity.Logging{
			UserId:  userId,
			Logging: message,
		})
		if err != nil {
			fmt.Println("insert log error")
		}
	}()
}
