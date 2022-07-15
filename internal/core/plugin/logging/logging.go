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

package logging

import (
	"github.com/acmestack/envcd/internal/pkg/constants"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/envcd/internal/pkg/plugin"
)

const (
	name = "logging"
)

type Logging struct {
	plugin.Plugin
}

func New() *Logging {
	l := &Logging{}
	l.Name = name
	l.Sort = constants.LoggingOrder
	return l
}

func (logging *Logging) Execute(context context.Context, chain executor.Chain) (ret interface{}, err error) {
	return chain.Execute(context)
}

func (logging *Logging) Skip(context context.Context) bool {
	return false
}
