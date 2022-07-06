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

// dictionary key value
type dictionary struct {
	size uint
	data map[interface{}]interface{}
}

// NewDictionary make new dictionary
func NewDictionary() *dictionary {
	return &dictionary{
		size: 0,
		data: make(map[any]any, 10),
	}
}

// Put new data to dictionary by key and value
func (dict *dictionary) Put(key interface{}, value interface{}) error {
	if dict == nil || dict.data == nil {
		return errors.New("the dictionary illegal state")
	}
	if dict.data[key] == nil {
		dict.size++
	}
	// if key is exist override or put it
	dict.data[key] = value
	return nil
}

// Get the data from dictionary by key
func (dict *dictionary) Get(key interface{}) (interface{}, error) {
	if dict == nil || dict.data == nil {
		return nil, errors.New("the dictionary illegal state")
	}
	return dict.data[key], nil
}

// Remove delete the data from dictionary by key
func (dict *dictionary) Remove(key interface{}) error {
	if dict == nil || dict.data == nil {
		return errors.New("the dictionary illegal state")
	}
	delete(dict.data, key)
	dict.size--
	return nil
}

// Size the dictionary data size
func (dict *dictionary) Size() uint {
	if dict == nil || dict.data == nil {
		return 0
	}
	return dict.size
}
