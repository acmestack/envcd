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

import "errors"

type Any = interface{}

// Dictionary key value
type Dictionary struct {
	size uint
	data map[interface{}]interface{}
}

// NewDictionary make new Dictionary
// todo with config store kind
func NewDictionary() *Dictionary {
	return &Dictionary{
		size: 0,
		data: make(map[Any]Any, 10),
	}
}

// Put new data to Dictionary by key and value
func (dict *Dictionary) Put(key interface{}, value interface{}) error {
	if dict == nil || dict.data == nil {
		return errors.New("the Dictionary illegal state")
	}
	if dict.data[key] == nil {
		dict.size++
	}
	// if key is exist override or put it
	dict.data[key] = value
	return nil
}

// Get the data from Dictionary by key
func (dict *Dictionary) Get(key interface{}) (interface{}, error) {
	if dict == nil || dict.data == nil {
		return nil, errors.New("the Dictionary illegal state")
	}
	return dict.data[key], nil
}

// Remove delete the data from Dictionary by key
func (dict *Dictionary) Remove(key interface{}) error {
	if dict == nil || dict.data == nil {
		return errors.New("the Dictionary illegal state")
	}
	delete(dict.data, key)
	dict.size--
	return nil
}

// Size the Dictionary data size
func (dict *Dictionary) Size() uint {
	if dict == nil || dict.data == nil {
		return 0
	}
	return dict.size
}
