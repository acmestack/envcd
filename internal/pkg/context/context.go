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

package context

import (
	"net/http"

	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/internal/pkg/result"
)

type EnvcdActionFunc func() *result.EnvcdResult

// Context for peer request
type Context struct {
	Uri         string
	Method      string
	Headers     http.Header
	ContentType string
	Cookies     []*http.Cookie
	Body        interface{}
	Request     *http.Request
	RequestId   string
	user        *entity.UserInfo
}

// AssignUser when the context's user is not assign
func (c *Context) AssignUser(user *entity.UserInfo) *Context {
	if c != nil && c.user == nil {
		c.user = user
	}
	return c
}

// User return the context assigned user
func (c *Context) User() *entity.UserInfo {
	if c == nil {
		return nil
	}
	return c.user
}
