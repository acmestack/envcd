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

package executor

import "github.com/acmestack/envcd/internal/pkg/context"

// Executor the executor
type Executor interface {

	// Execute execute code
	// Context come from every exector, data from dashboard
	//  @param context
	//  @param executor
	//  @return ret, error
	Execute(context context.Context, chain Chain) (ret interface{}, err error)

	// Skip current executor
	//  @return skip current executor or not
	Skip(context context.Context) bool

	// Sorted executor execute order
	//  @return order
	Sorted() uint8

	// Named executor name
	//  @return named for executor
	Named() string
}
