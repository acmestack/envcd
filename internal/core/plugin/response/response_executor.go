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

package response

import (
	"github.com/acmestack/envcd/internal/pkg/plugin"
)

type RespExecutor struct {
	responseExecutor *plugin.Executor
}

func (r RespExecutor) Execute(context interface{}, data interface{}, executor plugin.ExecutorChain) {
	//TODO implement me
	panic("implement me")
}

func (r RespExecutor) Skip(context interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (r RespExecutor) Order() uint8 {
	//TODO implement me
	panic("implement me")
}

func (r RespExecutor) Named() string {
	//TODO implement me
	panic("implement me")
}

// parse inner method, parse data from openapi
//  @param data data
//  @return result result
//  @return err error
func parse(data interface{}) (result map[string]string, err error) {
	//TODO parse data to map
	return nil, err
}
