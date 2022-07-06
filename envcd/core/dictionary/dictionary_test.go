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

package dictionary

import (
	"reflect"
	"testing"
)

func TestNewDictionary(t *testing.T) {
	tests := []struct {
		name string
		want *dictionary
	}{
		{
			want: NewDictionary(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDictionary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDictionary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dictionary_Get(t *testing.T) {
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
			dict := &dictionary{
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

func Test_dictionary_Put(t *testing.T) {
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
			dict := &dictionary{
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

func Test_dictionary_Remove(t *testing.T) {
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
			dict := &dictionary{
				size: tt.fields.size,
				data: tt.fields.data,
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

func Test_dictionary_Size(t *testing.T) {
	dict := NewDictionary()
	_ = dict.Put("a", "value")
	tests := []struct {
		name   string
		fields *dictionary
		want   uint
	}{
		{
			fields: dict,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := &dictionary{
				size: tt.fields.size,
				data: tt.fields.data,
			}
			if got := dict.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
