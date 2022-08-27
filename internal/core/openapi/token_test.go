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
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

// saltPassword Password generation Policy Test
func TestToken(t *testing.T) {
	tokenString, _ := generateToken(6, "userName")
	t.Logf(tokenString)
	token, _ := jwt.ParseWithClaims(tokenString, &authorizationClaims{}, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})
	if claim, ok := token.Claims.(*authorizationClaims); ok && token.Valid {
		t.Logf(claim.UserName)
	}
}
