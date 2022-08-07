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
	"reflect"
	"testing"

	"github.com/acmestack/godkits/gox/errorsx"
)

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

func TestInternalFailure(t *testing.T) {
	tests := []struct {
		name string
		want *EnvcdResult
	}{
		{
			want: InternalFailure(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InternalFailure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalFailure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternalFailureByError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{err: errorsx.Err("error")},
			want: InternalFailureByError(errorsx.Err("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InternalFailureByError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalFailureByError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailure(t *testing.T) {
	type args struct {
		envcdError envcdError
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{errorEnvcdInternalServerError},
			want: Failure0(errorEnvcdInternalServerError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failure0(tt.args.envcdError); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failure0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailureByError(t *testing.T) {
	type args struct {
		envcdError envcdError
		err        error
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{
				envcdError: errorEnvcdInternalServerError,
				err:        errorsx.Err("error"),
			},
			want: Failure(errorEnvcdInternalServerError, errorsx.Err("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failure(tt.args.envcdError, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_failure(t *testing.T) {
	type args struct {
		envcdErr envcdError
		err      error
	}
	tests := []struct {
		name string
		args args
		want *EnvcdResult
	}{
		{
			args: args{
				envcdErr: errorEnvcdInternalServerError,
				err:      nil,
			},
			want: failure(errorEnvcdInternalServerError, nil),
		},
		{
			args: args{
				envcdErr: errorEnvcdInternalServerError,
				err:      errorsx.Err("error"),
			},
			want: failure(errorEnvcdInternalServerError, errorsx.Err("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := failure(tt.args.envcdErr, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("failure() = %v, want %v", got, tt.want)
			}
		})
	}
}
