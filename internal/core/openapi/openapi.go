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
	"github.com/acmestack/envcd/internal/core/plugin"
	"github.com/acmestack/envcd/internal/core/plugin/logging"
	"github.com/acmestack/envcd/internal/core/plugin/permission"
	"github.com/acmestack/envcd/internal/core/storage"
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/envcd/pkg/entity/result"
	"github.com/acmestack/godkits/log"
	"github.com/gin-gonic/gin"
)

var requestIdHeader = "x-envcd-request-id"

type Openapi struct {
	exchange  *exchanger.Exchange
	storage   *storage.Storage
	executors []executor.Executor
	contexts  map[string]*context.Context
}

func Start(serverSetting *config.Server, exchange *exchanger.Exchange, storage *storage.Storage) {
	openapi := &Openapi{
		exchange:  exchange,
		storage:   storage,
		executors: []executor.Executor{logging.New(), permission.New()},
		contexts:  map[string]*context.Context{},
	}
	// sort plugin
	plugin.Sort(openapi.executors)
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

	log.Info("start http server listening %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Error("start http server error %v", err)
	}
}

// todo build Router
func (openapi *Openapi) buildRouter() *gin.Engine {
	router := gin.Default()
	// build context for peer request
	router.Use(openapi.buildContext)

	// login and logout
	router.POST("/login", openapi.login)
	router.GET("/logout", openapi.logout)

	// version 1 group
	v1 := router.Group("/v1")

	// user group routers
	usersGroup := v1.Group("/users")
	{
		// todo fuzzy => ?page=2&per_page=100&name=
		usersGroup.GET("", openapi.users)
		usersGroup.POST("", openapi.createUser)
		usersGroup.PUT("/:userId", openapi.updateUser)
		usersGroup.GET("/:userId", openapi.user)
		usersGroup.DELETE("/:userId", openapi.removeUser)

		// user's all scopespaces
		// todo fuzzy => ?page=2&per_page=100&scopespace_name=
		usersGroup.GET("/:userId/scopespaces", openapi.userScopeSpaces)
		// user's all dictionaries under one scopespace
		// todo fuzzy => ?page=2&per_page=100&scopespace_name=abc&dictionary_key=aaa
		usersGroup.GET("/:userId/scopespace/:scopeSpaceId/dictionaries", openapi.userDictionariesUnderScopeSpace)
		// user's all dictionaries
		// todo fuzzy => ?page=2&per_page=100&dictionary_key=aaa
		usersGroup.GET("/:userId/dictionaries", openapi.userDictionaries)
	}

	// scopespaces group routers
	scopeSpacesGroup := v1.Group("/scopespaces")
	{
		// todo fuzzy => ?page=2&per_page=100&userId=123&scopespace_name=
		scopeSpacesGroup.GET("", openapi.scopespaces)
		scopeSpacesGroup.POST("", openapi.createScopeSpace)
		scopeSpacesGroup.GET("/:scopeSpaceId", openapi.scopeSpace)
		scopeSpacesGroup.PUT("/:scopeSpaceId", openapi.updateScopeSpace)
		scopeSpacesGroup.DELETE("/:scopeSpaceId", openapi.removeScopeSpace)
	}

	// dictionaries group routers
	dictionariesGroup := v1.Group("/dictionaries")
	{
		// todo fuzzy => ?page=2&per_page=100&userId=123&name=
		dictionariesGroup.GET("", openapi.dictionaries)
		dictionariesGroup.POST("", openapi.createDictionary)
		dictionariesGroup.GET("/:dictId", openapi.dictionary)
		dictionariesGroup.PUT("/:dictId", openapi.updateDictionary)
		dictionariesGroup.DELETE("/:dictId", openapi.removeDictionary)
	}

	return router
}

// response to caller
func (openapi *Openapi) response(ginCtx *gin.Context, envcdAction context.EnvcdAction) {
	requestId := ginCtx.Request.Header.Get(requestIdHeader)
	c := openapi.contexts[requestId]
	ret := result.InternalServerErrorFailure(http.StatusText(http.StatusInternalServerError))
	if c != nil && c.RequestId == requestId {
		c.Action = envcdAction
		if exeRet := plugin.NewChain(openapi.executors).Execute(c); exeRet != nil {
			ret = exeRet
		}
	}
	ginCtx.JSON(ret.HttpStatusCode, ret.Data)
	delete(openapi.contexts, requestId)
}
