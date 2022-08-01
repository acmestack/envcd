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
	"io/ioutil"

	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/gin-gonic/gin"
)

// buildContext build plugin context
//  @param params params
//  @return *context.Context context
//  @return error error
func buildContext(ginCtx *gin.Context) (*context.Context, error) {
	ctx := &context.Context{
		Uri:         ginCtx.Request.RequestURI,
		Method:      ginCtx.Request.Method,
		Headers:     buildContextHeaders(ginCtx),
		ContentType: ginCtx.ContentType(),
		Cookies:     buildContextCookies(ginCtx),
		Body:        buildRequestBody(ginCtx),
		HttpRequest: ginCtx.Request,
	}
	if ctx != nil {
		return ctx, nil
	}
	return nil, nil
}

// parseContext parse context to envcd data
//  @param ctx context
//  @return *data.EnvcdData data
//  @return error error
func parseContext(ctx *context.Context) (*data.EnvcdData, error) {
	// TODO parse context to envcd data
	return nil, nil
}

// buildContextHeaders build plugin context headers
//  @param ginCtx gin context
//  @return map[string]interface{} ret
func buildContextHeaders(ginCtx *gin.Context) map[string]interface{} {
	maps := make(map[string]interface{})
	for k, v := range ginCtx.Request.Header {
		maps[k] = v
	}
	return maps
}

// buildContextCookies build plugin context cookies
//  @param ginCtx gin context
//  @return map[string]interface{} ret
func buildContextCookies(ginCtx *gin.Context) map[string]interface{} {
	maps := make(map[string]interface{})
	for k, v := range ginCtx.Request.Cookies() {
		maps[string(rune(k))] = v
	}
	return maps
}

// buildRequestBody build request body
//  @param ginCtx gin context
//  @return string request body
func buildRequestBody(ginCtx *gin.Context) string {
	all, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return ""
	}
	return string(all)
}
