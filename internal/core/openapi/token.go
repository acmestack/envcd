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
	"time"

	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/godkits/array"

	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/result"
	"github.com/golang-jwt/jwt/v4"
)

const (
	// hmacSecret secret
	hmacSecret = "9C035514A15F78"
	// tokenHeader
	tokenHeader = "token"
)

// authorizationClaims claims
type authorizationClaims struct {
	*jwt.RegisteredClaims
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

// convertTokenToUser parser token to user
func convertTokenToUser(tokenString string) *entity.UserInfo {
	token, err := jwt.ParseWithClaims(tokenString, &authorizationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})
	if err != nil {
		// todo log
		return nil
	}
	if claim, ok := token.Claims.(*authorizationClaims); ok && token.Valid {
		if claim.UserId == 0 {
			return nil
		}
		userInfo := &entity.UserInfo{}
		userInfo.Id = claim.UserId
		userInfo.Token = tokenString
		userInfo.Name = claim.UserName
		return userInfo
	}
	return nil
}

// generateToken with userId and userName
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
func (openapi *Openapi) validate(context *context.Context) context.EnvcdActionFunc {
	return func() *result.EnvcdResult {
		user := context.User()
		if user == nil {
			return result.Failure0(result.ErrorUserNotAuthorized)
		}
		param := entity.User{Id: user.Id}
		// query user by param
		users, err := dao.New(openapi.storage).SelectUser(param)
		if err != nil {
			return result.InternalFailure(err)
		}
		if array.Empty(users) || users[0].UserSession != user.Token {
			return result.Failure0(result.ErrorUserNotAuthorized)
		}
		return nil
	}
}
