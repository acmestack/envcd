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

package plugin

import (
	"reflect"
	"testing"

	"github.com/acmestack/envcd/internal/core/plugin/logging"
	"github.com/acmestack/envcd/internal/core/plugin/permission"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/envcd/pkg/entity/result"
	"github.com/acmestack/godkits/gox/errorsx"
)

func TestNewChain(t *testing.T) {
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
			want: NewChain([]executor.Executor{logging.New()}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChain(tt.args.executors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChain() = %v, want %v", got, tt.want)
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
		context *context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet interface{}
	}{
		{
			fields: fields{
				executors: []executor.Executor{logging.New()},
				index:     0,
			},
			args:    args{context: &context.Context{}},
			wantRet: result.Success(nil),
		},
		{
			fields: fields{
				executors: nil,
				index:     0,
			},
			args:    args{context: &context.Context{}},
			wantRet: result.InternalFailureByError(errorsx.Err("IIllegal state for plugin chain.")),
		},
		{
			fields: fields{
				executors: []executor.Executor{logging.New()},
				index:     0,
			},
			args:    args{context: &context.Context{Action: nil}},
			wantRet: result.Success(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executorChain := &Chain{
				executors: tt.fields.executors,
				index:     tt.fields.index,
			}
			gotRet := executorChain.Execute(tt.args.context)
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Execute() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		executors executorArray
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{executors: []executor.Executor{logging.New()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.executors)
		})
	}
}

func Test_executorArray_Len(t *testing.T) {
	tests := []struct {
		name string
		ea   executorArray
		want int
	}{
		{
			ea:   executorArray{logging.New()},
			want: 1,
		},
		{
			ea:   executorArray{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ea.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executorArray_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		ea   executorArray
		args args
		want bool
	}{
		{
			ea: executorArray{logging.New(), permission.New()},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
		{
			ea: executorArray{logging.New(), permission.New()},
			args: args{
				i: 1,
				j: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ea.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executorArray_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		ea   executorArray
		args args
	}{
		{
			ea: executorArray{logging.New(), permission.New()},
			args: args{
				i: 0,
				j: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ea.Swap(tt.args.i, tt.args.j)
		})
	}
}
