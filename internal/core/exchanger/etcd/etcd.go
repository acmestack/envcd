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
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/godkits/gox/stringsx"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

const (
	DefaultEtcdDialTimeout = "5"
)

type Etcd struct {
	ctx      context.Context
	client   *clientv3.Client
	endpoint string
}

// New make new etcd client
//  @param etcdConfig
//  @return *Etcd
func New(exchangerConnMetadata *config.ConnMetadata) *Etcd {

	ctx := context.Background()

	if exchangerConnMetadata.Type != "etcd" {
		log.Fatalf("Scheme is not eq = %v", exchangerConnMetadata.Type)
		return nil
	}

	endpoint := exchangerConnMetadata.Host + ":" + exchangerConnMetadata.Port
	if endpoint == "" {
		log.Fatalf("failed to get etcd url")
		return nil
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(stringsx.ToInt(DefaultEtcdDialTimeout)) * time.Second,
		Username:    exchangerConnMetadata.UserName,
		Password:    exchangerConnMetadata.Password,
	})

	if err != nil {
		log.Fatalf("failed to create etcd client %v", err)
		return nil
	}

	return &Etcd{
		ctx:      ctx,
		client:   cli,
		endpoint: endpoint,
	}
}

// Put put data into exchanger
//  @param key data identity
//  @param value data
func (etcd *Etcd) Put(key interface{}, value interface{}) error {
	cli := etcd.client
	putResponse, err := cli.Put(etcd.ctx, key.(string), value.(string), clientv3.WithPrevKV())
	if err != nil {
		log.Printf("failed put key/value [%s]/[%s],err: %v", key, value, err)
		return err
	}
	// if the key cover pre value, printf the pre value
	if putResponse.PrevKv != nil {
		log.Printf("Put etcd key = %s,pre value = %s", key, string(putResponse.PrevKv.Value))
		return nil
	}
	log.Printf("Put etcd key = %s, value = %s", key, value)
	return nil
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

// Remove remove data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Remove(key interface{}) error {
	cli := etcd.client
	delResponse, err := cli.Delete(etcd.ctx, key.(string), clientv3.WithPrevKV())

	if err != nil {
		log.Printf("failed delete key: %s,err: %v", key, err)
		return err
	}

	if delResponse.PrevKvs == nil {
		return nil
	}
	// printf the delete key/value
	for _, kvPair := range delResponse.PrevKvs {
		log.Printf("Delete key: %s,value: %s", kvPair.Key, kvPair.Value)
	}

	return nil
}
