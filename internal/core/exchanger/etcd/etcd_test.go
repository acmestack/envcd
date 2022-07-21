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
 * limit
 */

package etcd

import (
	"github.com/acmestack/envcd/internal/pkg/config"
	clientv3 "go.etcd.io/etcd/client/v3"

	"log"
	"testing"
)

var metadata = &config.Exchanger{
	Url: "",
	ConnMetadata: &config.ConnMetadata{
		Type:     "etcd",
		UserName: "root",
		Password: "root",
		Host:     "localhost:2379",
		Hostname: "localhost",
		Port:     2379,
	},
}

// Get get data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Get(key interface{}) (interface{}, error) {
	cli := etcd.client
	getResponse, err := cli.Get(etcd.ctx, key.(string))
	if err != nil {
		log.Printf("failed get value,err: %v", err)
		return "", err
	}

	if len(getResponse.Kvs) == 0 {
		log.Printf("key [%s] is not exist", key)
		return "", nil
	}

	value := getResponse.Kvs[0].Value
	return string(value), nil
}

// Find find data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Find(key interface{}) (interface{}, error) {
	cli := etcd.client
	rangeResponse, err := cli.Get(etcd.ctx, key.(string), clientv3.WithPrefix())
	result := make(map[string]string)
	if err != nil {
		log.Printf("failed get values,err: %v", err)
		return nil, err
	}
	length := len(rangeResponse.Kvs)
	if length == 0 {
		log.Printf("key [%s] is not exist", key)
		return result, nil
	}

	for _, resp := range rangeResponse.Kvs {
		result[string(resp.Key)] = string(resp.Value)
	}

	return result, nil
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(metadata),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
		})
	}
}

func TestEtcd_Put(t *testing.T) {
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(metadata),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			err := tt.want.Put("/test/a", "a")
			if err != nil {
				t.Errorf("failed to put, err = %v", err)
			}
			get, _ := tt.want.Get("/test/a")
			t.Logf("value = %s", get)
		})
	}
}

func TestEtcd_Get(t *testing.T) {
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(metadata),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			get, err := tt.want.Get("/test/a")
			if err != nil {
				t.Errorf("failed to get value,key = /test/a , err = %v", err)
			}
			t.Logf("value = %s", get)
		})
	}
}

func TestEtcd_Find(t *testing.T) {
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(metadata),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			values, err := tt.want.Find("/test/a")
			if err != nil {
				t.Errorf("failed to find key /test/a, err = %v", err)
			}
			t.Logf("values = %v", values)
		})
	}
}

func TestEtcd_Remove(t *testing.T) {
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(metadata),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			err := tt.want.Remove("/test/a")
			if err != nil {
				t.Errorf("failed to remove key /test/a, err = %v", err)
			}
		})
	}
}
