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
	"net/http"
	"reflect"
	"testing"
)

func TestInternalServerErrorFailure(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{message: "failure"},
			want: InternalServerErrorFailure("failure"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InternalServerErrorFailure(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalServerErrorFailure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuccess(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{data: "ok"},
			want: Success("ok"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Success(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Success() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey(t *testing.T) {
	CodeKey = "a"
	tests := []struct {
		name string
		want string
	}{
		{
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Success(nil); got == nil || (got.Data["a"] != successCode) {
				t.Errorf("Success() = %v", got)
			}
		})
	}
}

func TestFailure(t *testing.T) {
	type args struct {
		message        string
		httpStatusCode int
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{
				message:        http.StatusText(http.StatusBadRequest),
				httpStatusCode: http.StatusBadRequest,
			},
			want: Failure(http.StatusText(http.StatusBadRequest), http.StatusBadRequest),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failure(tt.args.message, tt.args.httpStatusCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailure1(t *testing.T) {
	type args struct {
		code           string
		message        string
		httpStatusCode int
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{
				code:           "ERROR",
				message:        http.StatusText(http.StatusBadRequest),
				httpStatusCode: http.StatusBadRequest,
			},
			want: Failure0("ERROR", http.StatusText(http.StatusBadRequest), http.StatusBadRequest),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failure0(tt.args.code, tt.args.message, tt.args.httpStatusCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failure0() = %v, want %v", got, tt.want)
			}
		})
	}
}
