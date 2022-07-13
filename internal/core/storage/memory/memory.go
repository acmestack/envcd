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

package memory

import "errors"

type Memory struct {
	size uint
	data map[interface{}]interface{}
}

func New() *Memory {
	return &Memory{
		size: 0,
		data: make(map[interface{}]interface{}, 10),
	}
}

func (memory *Memory) Put(key interface{}, value interface{}) error {
	if memory == nil || memory.data == nil {
		return errors.New("the illegal state of memory storage")
	}
	if memory.data[key] == nil {
		memory.size++
	}
	// if key is exist override or put it
	memory.data[key] = value
	return nil
}

func (memory *Memory) Get(key interface{}) (interface{}, error) {
	if memory == nil || memory.data == nil {
		return nil, errors.New("the illegal state of memory storage")
	}
	return memory.data[key], nil
}

func (memory *Memory) Find(key interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (memory *Memory) Remove(key interface{}) error {
	if memory == nil || memory.data == nil {
		return errors.New("the Dictionary illegal state")
	}
	delete(memory.data, key)
	memory.size--
	return nil
}
