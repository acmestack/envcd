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
	// user group
	userGroup := v1.Group("/user")
	{
		// TODO test
		userGroup.POST("", openapi.user)
		userGroup.PUT("/:id", openapi.updateUser)
		userGroup.GET("/:id", openapi.userById)
		userGroup.DELETE("/:id", openapi.removeUser)

		// TODO envcd scopespace
		userGroup.GET("/:userId/scopespace/:scopeSpaceId", openapi.scopeSpace)
		userGroup.POST("/:userId/scopespace", openapi.putScopeSpace)
		userGroup.PUT("/:userId/scopespace/:scopeSpaceId", openapi.updateScopeSpace)
		userGroup.DELETE("/:userId/scopespace/:scopeSpaceId", openapi.removeScopeSpace)

		// TODO envcd config
		userGroup.GET("/:userId/scopespace/:scopeSpaceId/dict/:dictId", openapi.dictionary)
		userGroup.POST("/:userId/scopespace/:scopeSpaceId/dict", openapi.putDictionary)
		userGroup.PUT("/:userId/scopespace/:scopeSpaceId/dict/:dictId", openapi.updateDictionary)
		userGroup.DELETE("/:userId/scopespace/:scopeSpaceId/dict/:dictId", openapi.removeDictionary)
	}
	envcdScopeSpace := v1.Group("/scopespace")
	{
		envcdScopeSpace.GET("/all")

	}
	envcdDictionary := v1.Group("/dictionary")
	{
		envcdDictionary.GET("/all")
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
