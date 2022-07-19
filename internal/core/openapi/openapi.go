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
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/godkits/log"
	"github.com/gin-gonic/gin"
)

type Openapi struct {
	exchange  *exchanger.Exchange
	storage   *storage.Storage
	executors []executor.Executor
}

func Start(serverSetting *config.Server, exchange *exchanger.Exchange, storage *storage.Storage) {
	openapi := &Openapi{
		exchange:  exchange,
		storage:   storage,
		executors: []executor.Executor{logging.New(), permission.New()},
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
		Addr:           fmt.Sprintf(":%d", serverSetting.HttpPort),
		Handler:        openapi.buildRouter(),
		ReadTimeout:    time.Duration(serverSetting.ReadTimeout),
		WriteTimeout:   time.Duration(serverSetting.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	log.Info("start http server listening %s", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Error("service error %v", err)
		return
	}
}

// todo build Router
func (openapi *Openapi) buildRouter() *gin.Engine {
	router := gin.Default()
	// router group
	adminGroup := router.Group("admin")
	{
		// TODO test
		adminGroup.GET("/login", openapi.login)
	}
	envcdGroup := router.Group("exchange")
	{
		envcdGroup.POST("/save", openapi.save)
	}
	return router
}
