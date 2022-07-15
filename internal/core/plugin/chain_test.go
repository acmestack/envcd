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

package plugin

import (
	"reflect"
	"testing"

	"github.com/acmestack/envcd/internal/core/plugin/logging"
	"github.com/acmestack/envcd/internal/pkg/executor"
)

func TestChain(t *testing.T) {
	type args struct {
		executors []executor.Executor
	}
	tests := []struct {
		name string
		args args
		want *Chain
	}{
		{
			args: args{executors: []executor.Executor{logging.New()}},
			want: New(logging.New()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.executors...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutorChain_Execute(t *testing.T) {
	type fields struct {
		executors []executor.Executor
		index     int
	}
	type args struct {
		context interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet interface{}
		wantErr bool
	}{
		{
			fields: fields{
				executors: []executor.Executor{logging.New()},
				index:     0,
			},
			args:    args{context: "string"},
			wantRet: "string",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executorChain := &Chain{
				executors: tt.fields.executors,
				index:     tt.fields.index,
			}
			gotRet, err := executorChain.Execute(tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Execute() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
