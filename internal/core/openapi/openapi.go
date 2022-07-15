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
	"github.com/acmestack/envcd/internal/core/plugin"
	"github.com/acmestack/envcd/internal/core/plugin/logging"
	"github.com/acmestack/envcd/internal/core/plugin/permission"
	"github.com/acmestack/envcd/internal/core/plugin/response"
	"github.com/acmestack/envcd/internal/core/storage"
	"github.com/acmestack/envcd/internal/envcd"
)

type Openapi struct {
	envcd       *envcd.Envcd
	storage     *storage.Storage
	pluginChain *plugin.Chain
}

func Start(envcd *envcd.Envcd, storage *storage.Storage) {
	openapi := &Openapi{
		envcd:   envcd,
		storage: storage,
		// the pluginChain for peer request router
		// todo sort plugin
		pluginChain: plugin.New(logging.New(), permission.New(), response.New()),
	}
	openapi.openRouter()
}

// todo open Router
func (openapi *Openapi) openRouter() {

}
