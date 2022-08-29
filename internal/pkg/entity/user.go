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

type UserInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Identity  int    `json:"identity"`
	State     string `json:"state"`
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type User struct {
	//TableName gobatis.TableName `user`
	Id          int       `column:"id"`
	Name        string    `column:"name"`
	Password    string    `column:"password"`
	Salt        string    `column:"salt"`
	Identity    int       `column:"identity"`
	State       string    `column:"state"`
	UserSession string    `column:"user_session"`
	CreatedAt   time.Time `column:"created_at"`
	UpdatedAt   time.Time `column:"updated_at"`
}

type PageUserParam struct {
	Page     int
	PageSize int
	Name     string
}
