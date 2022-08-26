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

	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/result"
	"github.com/golang-jwt/jwt/v4"
)

const (
	// hmacSecret secret
	hmacSecret = "9C035514A15F78"
)

// claims claims
type claims struct {
	*jwt.RegisteredClaims
	userId   int
	userName string
}

func generateToken(userId int, userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
		userId:   userId,
		userName: userName,
	})
	return token.SignedString([]byte(hmacSecret))
}

// validate current request
// user state and generateToken validation
func (openapi *Openapi) validate(context *context.Context) context.EnvcdActionFunc {
	return func() *result.EnvcdResult {
		if context.User == nil {
			return result.Failure0(result.ErrorUserNotAuthorized)
		}
		// todo validate generateToken
		return nil
	}
}
