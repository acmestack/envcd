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

import "net/http"

const (
	errorCodeUserNotFound          = "userNotFound"
	errorCodeUserExisted           = "userExisted"
	errorCodeCreateUser            = "userCreateFault"
	errorCodeUserPasswordIncorrect = "userPasswordIncorrect"
	errorCodeDictionaryNotExist    = "dictionaryNotExist"

	// errorCodeEnvcdInternalServerError internal server error code
	errorCodeEnvcdInternalServerError = "envcdInternalServerError"
	// envcdInternalServerErrorMessage internal server error message
	envcdInternalServerErrorMessage = "Envcd Internal Server Error, try again lately."
)

var (
	ErrorUserNotFound             envcdError
	ErrorUserExisted              envcdError
	ErrorCreateUser               envcdError
	ErrorUserPasswordIncorrect    envcdError
	ErrorDictionaryNotExist       envcdError
	errorEnvcdInternalServerError envcdError
)

func init() {
	ErrorCreateUser = envcdError{
		code:           errorCodeCreateUser,
		message:        "the user save error.",
		httpStatusCode: http.StatusOK,
	}
	ErrorUserNotFound = envcdError{
		code:           errorCodeUserNotFound,
		message:        "the user is not exist.",
		httpStatusCode: http.StatusBadRequest,
	}
	ErrorUserExisted = envcdError{
		code:           errorCodeUserExisted,
		message:        "the user is existed.",
		httpStatusCode: http.StatusOK,
	}
	ErrorUserPasswordIncorrect = envcdError{
		code:           errorCodeUserPasswordIncorrect,
		message:        "the password is incorrect for user.",
		httpStatusCode: http.StatusOK,
	}
	ErrorDictionaryNotExist = envcdError{
		code:           errorCodeDictionaryNotExist,
		message:        "the dictionary is not exist.",
		httpStatusCode: http.StatusBadRequest,
	}
	errorEnvcdInternalServerError = envcdError{
		code:           errorCodeEnvcdInternalServerError,
		message:        envcdInternalServerErrorMessage,
		httpStatusCode: http.StatusInternalServerError,
	}
}

type envcdError struct {
	code           string
	message        string
	httpStatusCode int
}
