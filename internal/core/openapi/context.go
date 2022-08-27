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
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// buildContext build plugin context
//  @param params params
//  @return *context.Context context
func (openapi *Openapi) buildContext(ginCtx *gin.Context) {
	// create request id by uuid
	requestId := core.RandomUUID()
	ginCtx.Request.Header.Add(requestIdHeader, requestId)
	openapi.contexts[requestId] = &context.Context{
		Uri:         ginCtx.Request.RequestURI,
		Method:      ginCtx.Request.Method,
		Headers:     ginCtx.Request.Header,
		ContentType: ginCtx.ContentType(),
		Cookies:     ginCtx.Request.Cookies(),
		Body:        ginCtx.Request.Body,
		Request:     ginCtx.Request,
		RequestId:   requestId,
		User:        openapi.parserUser(ginCtx),
	}
}

func (openapi *Openapi) parserUser(ginCtx *gin.Context) *entity.UserInfo {
	tokenString := ginCtx.GetHeader("token")
	if len(tokenString) == 0 {
		// not token
		return nil
	}
	token, _ := jwt.ParseWithClaims(tokenString, &authorizationClaims{}, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})
	if claim, ok := token.Claims.(*authorizationClaims); ok && token.Valid {
		if claim.UserId == 0 {
			return nil
		}
		userInfo := &entity.UserInfo{}
		userInfo.Id = claim.UserId
		userInfo.Token = tokenString
		return userInfo
	} else {
		return nil
	}
}
