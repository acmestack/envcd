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

func TestJson(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want EnvcdData
	}{
		{
			args: args{data: "{\"key\":\"value\"}"},
			want: Json("{\"key\":\"value\"}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Json(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Json() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperties(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want EnvcdData
	}{
		{
			args: args{data: "key=hello"},
			want: Properties("key=hello"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Properties(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Properties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want EnvcdData
	}{
		{
			args: args{data: "hello world"},
			want: String("hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToEnvcdData(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name string
		args args
		want EnvcdData
	}{
		{
			args: args{jsonString: "{\"format\":\"json\",\"data\":\"hello world\"}"},
			want: Json("hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToEnvcdData(tt.args.jsonString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToEnvcdData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToJson(t *testing.T) {
	type args struct {
		envcdData EnvcdData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{envcdData: String("hello world")},
			want: "{\"format\":\"string\",\"data\":\"hello world\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJson(tt.args.envcdData); got != tt.want {
				t.Errorf("ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYaml(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want EnvcdData
	}{
		{
			args: args{data: "yaml"},
			want: Yaml("yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yaml(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Yaml() = %v, want %v", got, tt.want)
			}
		})
	}
}
