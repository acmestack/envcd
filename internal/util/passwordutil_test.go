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

package util

import (
	"testing"
)

// EncryptPassword Password generation Policy Test
func TestEncryptPassword(t *testing.T) {
	password := "admin"
	salt := "07929137ab07437c933d6992321ef9fd"

	encryptPassword := EncryptPassword(password, salt)

	// login
	password = "admin"
	encryptPasswordLogin := EncryptPassword(password, salt)
	if encryptPasswordLogin != encryptPassword {
		t.Errorf("TestEncryptPassword() Login failed salt = %v, password %v", salt, password)
	}
}
