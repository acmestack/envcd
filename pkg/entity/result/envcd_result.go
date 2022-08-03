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

package result

import "net/http"

const (
	successCode = "SUCCESS"
	failureCode = "FAILURE"
)

var (
	ResultCodeKey    = "code"
	ResultMessageKey = "message"
	ResultDataKey    = "data"
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
	return &EnvcdResult{Data: map[string]interface{}{
		ResultCodeKey:    successCode,
		ResultMessageKey: "success",
		ResultDataKey:    data,
	}, HttpStatusCode: http.StatusOK}
}

// InternalServerErrorFailure response
//  @param message of error reason
//  @return *EnvcdResult
func InternalServerErrorFailure(message string) *EnvcdResult {
	return &EnvcdResult{Data: map[string]interface{}{
		ResultCodeKey:    failureCode,
		ResultMessageKey: message,
		ResultDataKey:    nil,
	}, HttpStatusCode: http.StatusInternalServerError}
}

// InternalServerErrorFailure response
//  @param message of error reason
//  @param httpStatusCode of response http status code
//  @return *EnvcdResult
func Failure(message string, httpStatusCode int) *EnvcdResult {
	return &EnvcdResult{Data: map[string]interface{}{
		ResultCodeKey:    failureCode,
		ResultMessageKey: message,
		ResultDataKey:    nil,
	}, HttpStatusCode: httpStatusCode}
}
