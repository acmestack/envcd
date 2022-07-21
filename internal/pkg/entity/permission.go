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

package entity

import (
	"github.com/acmestack/gobatis"
)

type Permission struct {
	PermissionTable gobatis.TableName "logging"
	*Base
	//Id        uint32
	UserId   uint32 `column:"user_id"`
	DataType uint8  `column:"data_type"`
	DataId   uint32 `column:"data_id"`
	Note     string `column:"note"`
	State    bool   `column:"state"`
	//createdAt time.Time
	//updatedAt time.Time
}
