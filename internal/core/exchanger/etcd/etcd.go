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

package etcd

import (
	"context"
	"log"
	"time"

	"github.com/acmestack/envcd/internal/pkg/config"
	"go.etcd.io/etcd/client/v3"
)

const (
	defaultEtcdDialTimeout = 5
	etcdExchangerType      = "etcd"
)

type Etcd struct {
	ctx    context.Context
	client *clientv3.Client
}

// New make new etcd client
//  @param etcdConfig
//  @return *Etcd
func New(exchanger *config.Exchanger) *Etcd {
	metadata := exchanger.ConnMetadata
	if metadata.Type != etcdExchangerType {
		log.Fatalf("invalid schema: %v", metadata.Type)
		return nil
	}
	endpoint := metadata.Host
	if endpoint == "" {
		log.Fatalf("failed to get etcd endpoint")
		return nil
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(defaultEtcdDialTimeout) * time.Second,
		Username:    metadata.UserName,
		Password:    metadata.Password,
	})

	if err != nil {
		log.Fatalf("failed to create etcd client %v", err)
		return nil
	}

	return &Etcd{
		ctx:    context.Background(),
		client: client,
	}
}

// Put put data into exchanger
//  @param key data identity
//  @param value data
func (etcd *Etcd) Put(key interface{}, value interface{}) error {
	response, err := etcd.client.Put(etcd.ctx, key.(string), value.(string), clientv3.WithPrevKV())
	if err != nil {
		log.Printf("failed put key/value [%s]/[%s],err: %v", key, value, err)
		return err
	}
	// if the key cover pre value, printf the pre value
	if response.PrevKv != nil {
		log.Printf("Put etcd key = %s,pre value = %s", key, string(response.PrevKv.Value))
		return nil
	}
	log.Printf("Put etcd key = %s, value = %s", key, value)
	return nil
}

// Remove remove data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Remove(key interface{}) error {
	response, err := etcd.client.Delete(etcd.ctx, key.(string), clientv3.WithPrevKV())
	if err != nil {
		log.Printf("failed delete key: %s,err: %v", key, err)
		return err
	}
	if response.PrevKvs == nil {
		return nil
	}
	// printf the delete key/value
	for _, kvPair := range response.PrevKvs {
		log.Printf("Delete key: %s,value: %s", kvPair.Key, kvPair.Value)
	}
	return nil
}
