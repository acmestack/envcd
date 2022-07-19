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

package openservice

import (
	"github.com/acmestack/envcd/internal/core/plugin"
	"github.com/acmestack/envcd/internal/core/plugin/logging"
	"github.com/acmestack/envcd/internal/core/plugin/permission"
	"github.com/acmestack/envcd/internal/core/storage"
	"github.com/acmestack/envcd/internal/envcd"
	"github.com/acmestack/envcd/internal/pkg/executor"
)

type OpenService struct {
	Envcd     *envcd.Envcd
	storage   *storage.Storage
	Executors []executor.Executor
}

func InitService(envcd *envcd.Envcd, storage *storage.Storage) *OpenService {
	openservice := &OpenService{
		Envcd:     envcd,
		storage:   storage,
		Executors: []executor.Executor{logging.New(), permission.New()},
	}
	// sort plugin
	plugin.Sort(openservice.Executors)
	return openservice
}
