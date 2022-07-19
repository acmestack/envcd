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
	routers2 "github.com/acmestack/envcd/internal/core/routers"
	openservice "github.com/acmestack/envcd/internal/core/service"
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/godkits/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Start(serverSetting *config.Server, openService *openservice.OpenService) {
	gin.SetMode(serverSetting.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverSetting.HttpPort),
		Handler:        routers2.InitRouter(openService),
		ReadTimeout:    time.Duration(serverSetting.ReadTimeout),
		WriteTimeout:   time.Duration(serverSetting.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	log.Info("[info] start http server listening %s", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Error("service error %v", err)
		return
	}
}
