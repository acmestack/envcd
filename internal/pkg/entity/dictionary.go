/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
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

import "time"

type Dictionary struct {
	//TableName gobatis.TableName `dictionary`
	Id           int       `column:"id" json:"id"`
	UserId       int       `column:"user_id" json:"userId"`
	ScopeSpaceId int       `column:"scopespace_id" json:"scopeSpaceId"`
	DictKey      string    `column:"dict_key" json:"dictKey"`
	DictValue    string    `column:"dict_value" json:"dictValue"`
	Version      string    `column:"version" json:"version"`
	State        string    `column:"state" json:"state"`
	CreatedAt    time.Time `column:"created_at" json:"createdAt"`
	UpdatedAt    time.Time `column:"updated_at" json:"updatedAt"`
}
