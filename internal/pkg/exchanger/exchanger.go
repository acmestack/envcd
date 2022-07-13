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

// Exchanger exchanger interface
type Exchanger interface {

	// Put put data into exchanger
	//  @param key data identity
	//  @param value data
	Put(key interface{}, value interface{}) error

	// Get get data from exchanger
	//  @param o data
	Get(key interface{}) (interface{}, error)

	// Find find data in exchanger
	//  @param o data
	Find(key interface{}) (interface{}, error)

	// Remove remove data from exchanger
	//  @param o data
	Remove(key interface{}) error
}
