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

var (
	ErrorUserNotFound             = aError("userNotFound", "the user is not exist.", http.StatusBadRequest)
	ErrorUserExisted              = aError("userExisted", "the user is existed.", http.StatusOK)
	ErrorCreateUser               = aError("userCreateFault", "the user save error.", http.StatusOK)
	ErrorUpdateUser               = aError("userUpdateFault", "the user update error.", http.StatusOK)
	ErrorUserPasswordIncorrect    = aError("userPasswordIncorrect", "the password is incorrect for user.", http.StatusOK)
	ErrorDictionaryNotExist       = aError("dictionaryNotExist", "the dictionary is not exist.", http.StatusBadRequest)
	ErrorEtcdPath                 = aError("etcd path error", "build etcd path error", http.StatusBadRequest)
	ErrorNotExistState            = aError("state not exist", "current state is error", http.StatusBadRequest)
	NilExchangePath               = aError("exchange path error", "exchange path is nil", http.StatusBadRequest)
	errorEnvcdInternalServerError = aError("envcdInternalServerError", "Envcd Internal Server Error, try again lately.", http.StatusInternalServerError)
)

type envcdError struct {
	code           string
	message        string
	httpStatusCode int
}

func aError(code string, message string, httpStatusCode int) envcdError {
	return envcdError{code: code, message: message, httpStatusCode: httpStatusCode}
}
