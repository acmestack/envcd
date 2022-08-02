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

package envcd

import (
	"github.com/acmestack/envcd/internal/core/exchanger"
	"github.com/acmestack/envcd/internal/core/openapi"
	"github.com/acmestack/envcd/internal/core/storage"
	"github.com/acmestack/envcd/internal/pkg/config"
)

// Start envcd by envcd exchangerConnMetadata config
//  @param exchangerConnMetadata the config for envcd
func Start(envcdConfig *config.Config) {
	// show start information & parser config
	envcdConfig.StartInformation()
	// start openapi with exchanger & storage
	openapi.Start(envcdConfig.Server,
		exchanger.Start(envcdConfig.Exchanger),
		storage.Start(envcdConfig.Storage))
}
