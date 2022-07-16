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

package plugin

import (
	"fmt"
	"sort"

	"github.com/acmestack/envcd/internal/core/plugin/response"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/godkits/gox/errorsx"
)

// executorArray for sort.Sort(Interface)
type executorArray []executor.Executor

func (ea executorArray) Len() int           { return len(ea) }
func (ea executorArray) Less(i, j int) bool { return ea[i].Sorted() < ea[j].Sorted() }
func (ea executorArray) Swap(i, j int)      { ea[i], ea[j] = ea[j], ea[i] }

// Chain the executor chain
// this is openapi chain, when http or client request into openapi, construct this chain
type Chain struct {
	executors []executor.Executor
	index     int
}

// Sort the plugins
func Sort(executors executorArray) {
	sort.Sort(executors)
}

// NewChain plugin chain for peer request
func NewChain(executors executorArray) *Chain {
	return &Chain{executors: executors, index: 0}
}

// Execute chain executor
//  @param context chain context
func (chain *Chain) Execute(context *context.Context) (ret interface{}, err error) {
	if chain == nil || chain.executors == nil || len(chain.executors) == 0 {
		message := "IIllegal state for plugin chain."
		return response.Failure(message), errorsx.Err(message)
	}
	if chain.index < len(chain.executors) {
		currentExecutor := chain.executors[chain.index]
		chain.index++
		if currentExecutor.Skip(context) {
			return chain.Execute(context)
		}
		// todo log
		fmt.Printf("plugin name '%v' sorted at '%v'\n", currentExecutor.Named(), currentExecutor.Sorted())
		// todo data
		return currentExecutor.Execute(context, chain)
	}
	return response.Success(nil), nil
}
