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
	"sort"

	"github.com/acmestack/envcd/internal/core/plugin/response"
	"github.com/acmestack/envcd/internal/pkg/executor"
	"github.com/acmestack/godkits/gox/errorsx"
)

type sortedExecutor []executor.Executor

func (s sortedExecutor) Len() int {
	return len(s)
}

func (s sortedExecutor) Less(i, j int) bool {
	return s[i].Order() < s[j].Order()
}

func (s sortedExecutor) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Chain the executor chain
// this is openapi chain, when http or client request into openapi, construct this chain
type Chain struct {
	executors []executor.Executor
	index     int
}

// Sort the plugins
func Sort(executors sortedExecutor) {
	sort.Sort(executors)
}

// NewChain plugin chain for peer request
func NewChain(executors sortedExecutor) *Chain {
	return &Chain{executors: executors, index: 0}
}

// Execute chain executor
//  @param context chain context
func (chain *Chain) Execute(context interface{}) (ret interface{}, err error) {
	if chain == nil || chain.executors == nil || len(chain.executors) == 0 {
		return nil, errorsx.Err("IIllegal state for plugin chain.")
	}
	if chain.index < len(chain.executors) {
		current := chain.executors[chain.index]
		chain.index++
		if current.Skip(context) {
			return chain.Execute(context)
		}
		// todo log
		// todo data
		return current.Execute(context, nil, chain)
	}
	return response.Success(nil), nil
}
