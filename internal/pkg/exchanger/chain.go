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

type Chain struct {
	elements []Exchanger
}

func New(exchangers ...Exchanger) *Chain {
	return &Chain{elements: exchangers}
}

func (chain *Chain) Put(key interface{}, value interface{}) error {
	for i := 0; i < len(chain.elements); i++ {
		err := chain.elements[i].Put(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (chain *Chain) Get(key interface{}) (interface{}, error) {
	for i := 0; i < len(chain.elements); i++ {
		data, err := chain.elements[i].Get(key)
		if err != nil {
			return nil, err
		}
		if data != nil {
			return data, nil
		}
	}
	return nil, nil
}

func (chain *Chain) Find(key interface{}) (interface{}, error) {
	for i := 0; i < len(chain.elements); i++ {
		data, err := chain.elements[i].Find(key)
		if err != nil {
			return nil, err
		}
		if data != nil {
			return data, nil
		}
	}
	return nil, nil
}

func (chain *Chain) Remove(key interface{}) error {
	for i := 0; i < len(chain.elements); i++ {
		err := chain.elements[i].Remove(key)
		if err != nil {
			return err
		}
	}
	return nil
}
