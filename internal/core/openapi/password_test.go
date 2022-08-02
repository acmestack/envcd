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

package openapi

import (
	"testing"
)

// saltPassword Password generation Policy Test
func TestPassword(t *testing.T) {
	plain := "admin"
	salt := "07929137ab07437c933d6992321ef9fd"

	password := saltPassword(plain, salt)

	// login
	plain = "admin"
	if saltPassword(plain, salt) != password {
		t.Errorf("TestPassword() Login failed salt = %v, saltPassword %v", salt, plain)
	}
}
