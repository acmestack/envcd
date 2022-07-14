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

type ExchangeChain struct {
	exchangers []Exchanger
}

func Chain(exchangers ...Exchanger) *ExchangeChain {
	return &ExchangeChain{exchangers: exchangers}
}

func (chain *ExchangeChain) Put(key interface{}, value interface{}) error {
	for i := 0; i < len(chain.exchangers); i++ {
		if err := chain.exchangers[i].Put(key, value); err != nil {
			return err
		}
		if err := chain.exchangers[i].Sync(key, value); err != nil {
			return err
		}
	}
	return nil
}

func (chain *ExchangeChain) Get(key interface{}) (interface{}, error) {
	for i := 0; i < len(chain.exchangers); i++ {
		data, err := chain.exchangers[i].Get(key)
		if err != nil {
			return nil, err
		}
		if data != nil {
			return data, nil
		}
	}
	return nil, nil
}

func (chain *ExchangeChain) Find(key interface{}) (interface{}, error) {
	for i := 0; i < len(chain.exchangers); i++ {
		data, err := chain.exchangers[i].Find(key)
		if err != nil {
			return nil, err
		}
		if data != nil {
			return data, nil
		}
	}
	return nil, nil
}

func (chain *ExchangeChain) Remove(key interface{}) error {
	for i := 0; i < len(chain.exchangers); i++ {
		if err := chain.exchangers[i].Remove(key); err != nil {
			return err
		}
	}
	return nil
}
