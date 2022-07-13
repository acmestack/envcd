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
	"errors"

	"github.com/acmestack/envcd/internal/core/storage/memory"
	"github.com/acmestack/envcd/internal/pkg/storage"
)

type Any = interface{}

// Dictionary key value
type Dictionary struct {
	storage storage.Storage
}

// NewDictionary make new Dictionary
// todo with config storage kind
func NewDictionary() *Dictionary {
	return &Dictionary{
		storage: memory.New(),
	}
}

// Put new data to Dictionary by key and value
func (dict *Dictionary) Put(key interface{}, value interface{}) error {
	if dict == nil || dict.storage == nil {
		return errors.New("the Dictionary illegal state")
	}
	return dict.storage.Put(key, value)
}

// Get the data from Dictionary by key
func (dict *Dictionary) Get(key interface{}) (interface{}, error) {
	if dict == nil || dict.storage == nil {
		return nil, errors.New("the Dictionary illegal state")
	}
	return dict.storage.Get(key)
}

// Find delete the data from Dictionary by key
func (dict *Dictionary) Find(key interface{}) (interface{}, error) {
	if dict == nil || dict.storage == nil {
		return nil, errors.New("the Dictionary illegal state")
	}
	return dict.storage.Find(key)
}

// Remove delete the data from Dictionary by key
func (dict *Dictionary) Remove(key interface{}) error {
	if dict == nil || dict.storage == nil {
		return errors.New("the Dictionary illegal state")
	}
	return dict.storage.Remove(key)
}
