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
	"errors"

	"github.com/acmestack/envcd/internal/core/exchanger/etcd"
	"github.com/acmestack/envcd/internal/pkg/exchanger"
)

type Exchanger struct {
	exchanger exchanger.Exchanger
}

// Start the Exchanger
func Start() *Exchanger {
	return &Exchanger{
		exchanger: etcd.New(),
	}
}

// Put new data to Exchanger by key and value
func (exchanger *Exchanger) Put(key interface{}, value interface{}) error {
	if exchanger == nil || exchanger.exchanger == nil {
		return errors.New("IIllegal state for exchanger")
	}
	return exchanger.exchanger.Put(key, value)
}

// Get the data from Exchanger by key
func (exchanger *Exchanger) Get(key interface{}) (interface{}, error) {
	if exchanger == nil || exchanger.exchanger == nil {
		return nil, errors.New("IIllegal state for exchanger")
	}
	return exchanger.exchanger.Get(key)
}

// Find delete the data from Exchanger by key
func (exchanger *Exchanger) Find(key interface{}) (interface{}, error) {
	if exchanger == nil || exchanger.exchanger == nil {
		return nil, errors.New("IIllegal state for exchanger")
	}
	return exchanger.exchanger.Find(key)
}

// Remove delete the data from Exchanger by key
func (exchanger *Exchanger) Remove(key interface{}) error {
	if exchanger == nil || exchanger.exchanger == nil {
		return errors.New("IIllegal state for exchanger")
	}
	return exchanger.exchanger.Remove(key)
}
