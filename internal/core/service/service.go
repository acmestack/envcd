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

package service

import (
	"fmt"
	"github.com/acmestack/envcd/internal/core/service/routers"
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/godkits/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Start(envcdConfig *config.Config) {
	gin.SetMode(envcdConfig.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := envcdConfig.ServerSetting.ReadTimeout
	writeTimeout := envcdConfig.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", envcdConfig.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    time.Duration(readTimeout),
		WriteTimeout:   time.Duration(writeTimeout),
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Info("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Error("service error %v", err)
		return
	}
}
