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
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/acmestack/godkits/gox/encodingx/jsonx"
	"github.com/acmestack/godkits/gox/stringsx"
	"strings"
	"sync"
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
	// TODO write log
	// FIXME log.info(printLog(context))
	return chain.Execute(context)
}

func (logging *Logging) Skip(context *context.Context) bool {
	return false
}

func printLog(ctx *context.Context) (ret string) {
	var wg sync.WaitGroup
	wg.Add(1)
	builder := strings.Builder{}
	build := &stringsx.Builder{Builder: builder}
	_, _ = build.JoinString("\n[Request Information]\n")
	_, _ = build.JoinString("Request Uri:", ctx.Uri, "\n")
	_, _ = build.JoinString("Request Method:", ctx.Method, "\n")
	header, _ := jsonx.ToJsonString(ctx.Headers)
	_, _ = build.JoinString("Request Headers:", header, "\n")
	_, _ = build.JoinString("Request ContentType:", ctx.ContentType, "\n")
	params, _ := jsonx.ToJsonString(ctx.Parameters)
	_, _ = build.JoinString("Request Parameters:", params, "\n\n")
	_, _ = build.JoinString("[Request Body]\n")
	requestBody, _ := jsonx.ToJsonString(ctx.Body)
	_, _ = build.JoinString("Request Boyd:", requestBody, "\n\n")
	_, _ = build.JoinString("[Response Body]\n")

	//TODO chan accept action complete and return response
	responseBody, _ := jsonx.ToJsonString(ctx.Action)
	_, _ = build.JoinString("Response Body:", responseBody, "\n")

	wg.Done()
	return build.String()
}
