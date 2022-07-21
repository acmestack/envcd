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

package config

import (
	"github.com/acmestack/godkits/gox/stringsx"
	"reflect"
	"testing"
)

func TestConnMetadata_information(t *testing.T) {
	type fields struct {
		Type     string
		UserName string
		Password string
		Host     string
		Port     string
	}
	type args struct {
		t string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				Type:     "etcd",
				UserName: "user",
				Password: "x",
				Host:     "localhost",
				Port:     "2379",
			},
			args: args{
				t: "someone",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connMedata := &ConnMetadata{
				Type:     tt.fields.Type,
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				Hostname: tt.fields.Host,
				Port:     stringsx.ToInt(tt.fields.Port),
			}
			connMedata.information(tt.args.t)
		})
	}
}

func Test_parser(t *testing.T) {
	type args struct {
		connection string
	}
	tests := []struct {
		name string
		args args
		want *ConnMetadata
	}{
		{
			args: args{connection: "etcd://user:password@localhost:2379"},
			want: &ConnMetadata{
				Type:     "etcd",
				UserName: "user",
				Password: "password",
				Host:     "localhost:2379",
				Hostname: "localhost",
				Port:     2379,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parser(tt.args.connection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser() = %v, want %v", got, tt.want)
			}
		})
	}
}
