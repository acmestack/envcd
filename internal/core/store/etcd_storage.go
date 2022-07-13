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

package store

import "github.com/acmestack/envcd/internal/core/base"

// todo memory storage
// todo write log first, write data second
// todo log module, chain

type EtcdStorage struct {
	Storage
}

type EtcdPlugin struct {
	base.PluginExecutor
}

// Put put data into etcd
//  @receiver store etcd store
//  @param o data
func (store EtcdStorage) Put(o interface{}) {

}

// Get get data from etcd
//  @receiver store etcd store
//  @param o data
func (store EtcdStorage) Get(o interface{}) {

}

// Find find data from etcd
//  @receiver store etcd store
//  @param o data
func (store EtcdStorage) Find(o interface{}) {

}

// Remove remove data from etcd
//  @receiver store etcd store
//  @param o data
func (store EtcdStorage) Remove(o interface{}) {

}

func (plugin EtcdPlugin) Execute(context interface{}, data interface{}, executor base.PluginChainExecutor) {

}

func (plugin EtcdPlugin) Skip(context interface{}) bool {
	return false
}
