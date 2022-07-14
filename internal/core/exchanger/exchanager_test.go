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

package exchanger

import (
	"reflect"
	"testing"

	"github.com/acmestack/envcd/internal/core/storage/memory"
	"github.com/acmestack/envcd/internal/pkg/exchanger"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Exchanger
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Exchanger_Get(t *testing.T) {
	type fields struct {
		chain *exchanger.ExchangeChain
	}
	type args struct {
		key interface{}
	}
	mem := memory.New()
	_ = mem.Put("a", "value")
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{
				chain: exchanger.Chain(mem),
			},
			args:    args{key: "a"},
			want:    "value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Exchanger{
				chain: tt.fields.chain,
			}
			got, err := dict.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Exchanger_Put(t *testing.T) {
	type fields struct {
		chain *exchanger.ExchangeChain
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{

			fields: fields{
				chain: exchanger.Chain(memory.New()),
			},
			args:    args{key: "a", value: "value"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Exchanger{
				chain: tt.fields.chain,
			}
			if err := dict.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Exchanger_Remove(t *testing.T) {
	type fields struct {
		chain *exchanger.ExchangeChain
	}
	type args struct {
		key interface{}
	}
	mem := memory.New()
	_ = mem.Put("a", "value")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{
				chain: exchanger.Chain(mem),
			},
			args:    args{key: "a"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Exchanger{
				chain: tt.fields.chain,
			}
			if err := dict.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if data, err := dict.Get(tt.args.key); (err != nil) || data == "value" {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
