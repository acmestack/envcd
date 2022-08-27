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
	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/result"
	"github.com/golang-jwt/jwt/v4"
)

const (
	// hmacSecret secret
	hmacSecret = "9C035514A15F78"
)

// claims claims
type authorizationClaims struct {
	*jwt.RegisteredClaims
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

func generateToken(userId int, userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &authorizationClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
		UserId:   userId,
		UserName: userName,
	})
	return token.SignedString([]byte(hmacSecret))
}

// validate current request
// user state and generateToken validation
func (openapi *Openapi) validate(context *context.Context, ctx *gin.Context) context.EnvcdActionFunc {
	return func() *result.EnvcdResult {
		uri := ctx.Request.RequestURI
		if !strings.Contains(uri, "login") {
			return nil
		}
		tokenString := ctx.GetHeader("token")
		if len(tokenString) == 0 {
			// not token
			return result.Failure0(result.ErrorUserNotAuthorized)
		}
		token, _ := jwt.ParseWithClaims(tokenString, &authorizationClaims{}, func(token *jwt.Token) (interface{}, error) {
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(hmacSecret), nil
		})
		if claim, ok := token.Claims.(*authorizationClaims); ok && token.Valid {
			if claim.UserId == 0 {
				return result.Failure0(result.ErrorUserNotAuthorized)
			}
			param := entity.User{Id: claim.UserId}
			// query user by param
			users, _ := dao.New(openapi.storage).SelectUser(param)
			if len(users) == 0 {
				return result.Failure0(result.ErrorUserNotAuthorized)
			}
			if users[0].UserSession != tokenString {
				return result.Failure0(result.ErrorUserNotAuthorized)
			}
			userInfo := &entity.UserInfo{}
			userInfo.Id = users[0].Id
			userInfo.Name = users[0].Name
			userInfo.Identity = users[0].Identity
			userInfo.State = users[0].State
			userInfo.CreatedAt = users[0].CreatedAt.Format("2006-01-02 15:04:05")
			userInfo.UpdatedAt = users[0].UpdatedAt.Format("2006-01-02 15:04:05")
			context.User = userInfo
		} else {
			return result.Failure0(result.ErrorUserNotAuthorized)
		}
		if context.User == nil {
			return result.Failure0(result.ErrorUserNotAuthorized)
		}
		return nil
	}
}
