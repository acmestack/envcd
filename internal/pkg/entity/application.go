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

type Application struct {
	// ApplicationTable declare application table
	ApplicationTable gobatis.TableName "application"
	*Base
	//Id        uint32    `column:"id"`
	Name  string `column:"name"`
	Note  string `column:"note"`
	state bool   `column:"state"`
	//createdAt time.Time `column:"created_at"`
	//updatedAt time.Time `column:"updated_at"`
}
