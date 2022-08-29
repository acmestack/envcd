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

package openapi

import (
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/godkits/core"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-gonic/gin"
)

const (
	requestIdHeader = "x-envcd-request-id"
	// tokenHeader
	tokenHeader = "x-envcd-token"
)

// buildContext build plugin context
//  @param params params
//  @return *context.Context context
func (openapi *Openapi) buildContext(ginCtx *gin.Context) {
	// create request id by uuid
	requestId := core.RandomUUID()
	ginCtx.Request.Header.Add(requestIdHeader, requestId)
	openapi.contexts[requestId] = (&context.Context{
		Uri:         ginCtx.Request.RequestURI,
		Method:      ginCtx.Request.Method,
		Headers:     ginCtx.Request.Header,
		ContentType: ginCtx.ContentType(),
		Cookies:     ginCtx.Request.Cookies(),
		Body:        ginCtx.Request.Body,
		Request:     ginCtx.Request,
		RequestId:   requestId,
	}).AssignUser(openapi.parserUser(ginCtx))
}

func (openapi *Openapi) parserUser(ginCtx *gin.Context) *entity.UserInfo {
	tokenString := ginCtx.GetHeader(tokenHeader)
	if stringsx.Empty(tokenString) {
		// not token
		return nil
	}
	return convertTokenToUser(tokenString)
}
