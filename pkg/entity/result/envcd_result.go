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

package result

import (
	"fmt"
	"net/http"
)

const (
	successCode = "success"
)

var (
	CodeKey    = "code"
	MessageKey = "message"
	DataKey    = "data"
)

// EnvcdResult for response
type EnvcdResult struct {
	// response data
	Data map[string]interface{}
	// response http status code
	HttpStatusCode int
}

// Success response
//  @param data
//  @return *EnvcdResult
func Success(data interface{}) *EnvcdResult {
	return &EnvcdResult{
		Data: map[string]interface{}{
			CodeKey:    successCode,
			MessageKey: "success",
			DataKey:    data,
		},
		HttpStatusCode: http.StatusOK,
	}
}

// InternalFailure response
//  @return *EnvcdResult
func InternalFailure() *EnvcdResult {
	return failure(errorEnvcdInternalServerError, nil)
}

// InternalFailureByError response
//  @param message of error format: envcd internal error message : detail error [code]
//  @return *EnvcdResult
func InternalFailureByError(err error) *EnvcdResult {
	return failure(errorEnvcdInternalServerError, err)
}

// Failure response
//  @param envcdError of envcd error
//  @param err of handler error
//  @return *EnvcdResult
func Failure(envcdError envcdError, err error) *EnvcdResult {
	return failure(envcdError, err)
}

// Failure0 response
//  @param envcdError of envcd error
//  @return *EnvcdResult
func Failure0(envcdErr envcdError) *EnvcdResult {
	return failure(envcdErr, nil)
}

// failure response
//  @param envcdError of envcd error
//  @param err of error instance
//  @return *EnvcdResult
func failure(envcdErr envcdError, err error) *EnvcdResult {
	message := envcdErr.message
	envcdErr.message = fmt.Sprintf("%s [%s]", message, envcdErr.code)
	if err != nil && err.Error() != "" {
		envcdErr.message = fmt.Sprintf("%s [%s: %s]", message, envcdErr.code, err.Error())
	}
	return &EnvcdResult{
		Data: map[string]interface{}{
			CodeKey:    envcdErr.code,
			MessageKey: envcdErr.message,
		},
		HttpStatusCode: envcdErr.httpStatusCode,
	}
}
