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

package main

import (
	"flag"
	"github.com/acmestack/envcd/internal/core/openapi"
	"github.com/acmestack/envcd/internal/core/storage"
	"github.com/acmestack/envcd/internal/envcd"
	"github.com/acmestack/envcd/internal/pkg/config"
)

func main() {
	configFile := flag.String("config", "config/envcd.yaml", "envcd -config config/envcd.yaml")
	flag.Parse()
	configData := config.NewConfig(configFile)
	// show start information & parser config
	configData.StartInformation()
	// start openapi with exchanger & storage
	openapi.Start(configData.ServerSetting,
		envcd.Start(configData.ExchangerConnMetadata),
		storage.Start(configData.MysqlConnMetadata))
}
