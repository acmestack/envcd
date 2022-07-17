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

package envcd

import (
	"errors"
	"github.com/acmestack/envcd/internal/pkg/config"
	"reflect"
	"testing"

	"github.com/acmestack/envcd/internal/core/exchanger/etcd"
	"github.com/acmestack/envcd/internal/pkg/exchanger"
)

type Memory struct {
	size uint
	data map[interface{}]interface{}
}

func New() *Memory {
	return &Memory{
		size: 0,
		data: make(map[interface{}]interface{}, 10),
	}
}

func (memory *Memory) Put(key interface{}, value interface{}) error {
	if memory == nil || memory.data == nil {
		return errors.New("the illegal state of memory exchanger")
	}
	if memory.data[key] == nil {
		memory.size++
	}
	// if key is exist override or put it
	memory.data[key] = value
	return nil
}

func (memory *Memory) Get(key interface{}) (interface{}, error) {
	if memory == nil || memory.data == nil {
		return nil, errors.New("the illegal state of memory exchanger")
	}
	return memory.data[key], nil
}

func (memory *Memory) Find(key interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (memory *Memory) Remove(key interface{}) error {
	if memory == nil || memory.data == nil {
		return errors.New("the Dictionary illegal state")
	}
	delete(memory.data, key)
	memory.size--
	return nil
}
func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Memory
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Memory_Get(t *testing.T) {
	type fields struct {
		size uint
		data map[interface{}]interface{}
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{
				size: 1,
				data: map[interface{}]interface{}{
					"a": "value",
				},
			},
			args:    args{key: "a"},
			want:    "value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Memory{
				size: tt.fields.size,
				data: tt.fields.data,
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

func Test_Memory_Put(t *testing.T) {
	type fields struct {
		size uint
		data map[interface{}]interface{}
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
				size: 1,
				data: map[interface{}]interface{}{},
			},
			args:    args{key: "a", value: "value"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Memory{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if err := dict.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if data, err := dict.Get(tt.args.key); (err != nil) || data != tt.args.value {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Memory_Remove(t *testing.T) {
	type fields struct {
		size uint
		data map[interface{}]interface{}
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{

			fields: fields{
				size: 1,
				data: map[interface{}]interface{}{
					"a": "value",
				},
			},
			args:    args{key: "a"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Memory{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if err := dict.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if data, err := dict.Get(tt.args.key); ((err != nil) != tt.wantErr) && data != "" {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStart(t *testing.T) {
	exchangerConnMetadata := &config.ConnMetadata{
			Type:     "etcd",
			UserName: "root",
			Password: "root",
			Host:     "localhost",
			Port:     "2379",
	}

	e := &Envcd{exchanger: etcd.New(exchangerConnMetadata)}
	tests := []struct {
		name string
		want *Envcd
	}{
		{
			want: e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := e; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Exchanger_Get(t *testing.T) {
	type fields struct {
		exchanger exchanger.Exchanger
	}
	type args struct {
		key interface{}
	}
	mem := New()
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
				exchanger: mem,
			},
			args:    args{key: "a"},
			want:    "value",
			wantErr: false,
		},
		{
			fields: fields{
				exchanger: nil,
			},
			args:    args{key: "a"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Envcd{
				exchanger: tt.fields.exchanger,
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
		exchanger exchanger.Exchanger
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
				exchanger: New(),
			},
			args:    args{key: "a", value: "value"},
			wantErr: false,
		},
		{

			fields: fields{
				exchanger: nil,
			},
			args:    args{key: "a", value: "value"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Envcd{
				exchanger: tt.fields.exchanger,
			}
			if err := dict.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Exchanger_Remove(t *testing.T) {
	type fields struct {
		exchanger exchanger.Exchanger
	}
	type args struct {
		key interface{}
	}
	mem := New()
	_ = mem.Put("a", "value")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{
				exchanger: mem,
			},
			args:    args{key: "a"},
			wantErr: false,
		},
		{
			fields: fields{
				exchanger: nil,
			},
			args:    args{key: "a"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &Envcd{
				exchanger: tt.fields.exchanger,
			}
			if err := dict.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if data, err := dict.Get(tt.args.key); ((err != nil) != tt.wantErr) && data != "" {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
