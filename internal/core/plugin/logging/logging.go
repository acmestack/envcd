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

package logging

import (
	"github.com/acmestack/envcd/internal/pkg/constants"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/envcd/internal/pkg/plugin"
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/acmestack/godkits/gox/encodingx/jsonx"
	"github.com/acmestack/godkits/gox/stringsx"
)

const (
	name = "Logging"
)

type Logging struct {
	plugin.Plugin
}

func New() *Logging {
	l := &Logging{}
	l.Name = name
	l.Sort = constants.LoggingSorted
	return l
}

func (logging *Logging) Execute(context *context.Context, chain executor.Chain) (*data.EnvcdResult, error) {
	printLog(context)
	return chain.Execute(context)
}

func (logging *Logging) Skip(context *context.Context) bool {
	return false
}

func printLog(ctx *context.Context) {
	build := stringsx.Builder{}
	header, _ := jsonx.ToJsonString(ctx.Headers)
	params, _ := jsonx.ToJsonString(ctx.Parameters)
	requestBody, _ := jsonx.ToJsonString(ctx.Body)
	responseBody, _ := jsonx.ToJsonString(ctx.Action)
	_, err := build.JoinString("\n[Request Information]\n",
		"Request Uri:", ctx.Uri, "\n",
		"Request Method:", ctx.Method, "\n",
		"Request Headers:", header, "\n",
		"Request ContentType:", ctx.ContentType, "\n",
		"Request Parameters:", params, "\n\n",
		"[Request Body]\n",
		"Request Boyd:", requestBody, "\n\n",
		"[Response Body]\n",
		"Response Body:", responseBody, "\n",
	)
	if err == nil {
		s := build.String()
		println(s)
	}
}
