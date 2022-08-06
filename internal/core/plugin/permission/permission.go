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

package permission

import (
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/envcd/internal/pkg/plugin"
	"github.com/acmestack/envcd/pkg/entity/result"
)

const (
	name = "Permission"
)

type Permission struct {
	plugin.Plugin
}

func New() *Permission {
	p := &Permission{}
	p.Name = name
	p.Sort = plugin.PermissionSorted
	return p
}

func (permission *Permission) Execute(context *context.Context, chain executor.Chain) *result.EnvcdResult {
	if ret := permission.tokenValidate(context); ret != nil {
		return ret
	}
	if context.PermissionAction != nil {
		if ret := context.PermissionAction(); ret != nil {
			return ret
		}
	}
	return chain.Execute(context)
}

func (permission Permission) tokenValidate(context *context.Context) *result.EnvcdResult {
	// todo token validate
	return nil
}

func (permission *Permission) Skip(context *context.Context) bool {
	return false
}
