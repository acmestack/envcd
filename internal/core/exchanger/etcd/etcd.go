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

// todo memory exchanger
// todo write log first, write data second
// todo log module, chain

type Etcd struct {
}

func New() *Etcd {
	return &Etcd{}
}

// Put put data into exchanger
//  @param key data identity
//  @param value data
func (etcd *Etcd) Put(key interface{}, value interface{}) error {
	//TODO implement me
	panic("implement me")
}

// Get get data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Get(key interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

// Find find data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Find(key interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

// Remove remove data from etcd
//  @receiver exchanger etcd exchanger
//  @param o data
func (etcd *Etcd) Remove(key interface{}) error {
	//TODO implement me
	panic("implement me")
}
