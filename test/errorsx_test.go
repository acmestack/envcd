/*
 * Copyright (c) 2022, OpeningO
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

package test

import (
	"reflect"
	"testing"

	"github.com/openingo/godkits/gox/errorsx"
)

func TestErr(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args:    args{message: "error message"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := errorsx.Err(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Err() error = %v, wantErr %v, logError %v", err, tt.wantErr, errorsx.LogError(err))
			}
		})
	}
}

func TestLogError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr any
	}{
		{
			args:    args{err: nil},
			wantErr: nil,
		},
		{
			args:    args{err: errorsx.Err("error")},
			wantErr: errorsx.Err("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := errorsx.LogError(tt.args.err); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("LogError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
