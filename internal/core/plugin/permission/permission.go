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

package permission

import (
	"github.com/acmestack/envcd/internal/pkg/executor"
)

type Permission struct{}

func New() *Permission {
	return &Permission{}
}

func (permission *Permission) Execute(context interface{}, data interface{}, chain executor.Chain) (ret interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (permission *Permission) Skip(context interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (permission *Permission) Order() uint8 {
	//TODO implement me
	panic("implement me")
}

func (permission *Permission) Named() string {
	//TODO implement me
	panic("implement me")
}
