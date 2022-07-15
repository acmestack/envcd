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

package response

const (
	successCode = "SUCCESS"
	failureCode = "FAILURE"
)

var (
	CodeKey    = "code"
	MessageKey = "message"
	DataKey    = "data"
)

// Data for response
type Data struct {
	codeKey    string
	messageKey string
	dataKey    string
	Data       map[interface{}]interface{}
}

// Success response
//  @param data
//  @return *Data
func Success(data interface{}) *Data {
	return &Data{Data: map[interface{}]interface{}{
		CodeKey:    successCode,
		MessageKey: "success",
		DataKey:    data,
	}}
}

// Failure response
//  @param message of error reason
//  @return *Data
func Failure(message string) *Data {
	return &Data{Data: map[interface{}]interface{}{
		CodeKey:    failureCode,
		MessageKey: message,
		DataKey:    nil,
	}}
}
