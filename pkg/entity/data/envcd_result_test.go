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

package data

import (
	"reflect"
	"testing"
)

func TestFailure(t *testing.T) {
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
			want: Failure("failure"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failure(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failure() = %v, want %v", got, tt.want)
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
	ResultCodeKey = "a"
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
