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
	"github.com/acmestack/godkits/core"
	"github.com/acmestack/godkits/gox/cryptox/md5x"
	"github.com/acmestack/godkits/gox/cryptox/sha256x"
)

// saltPassword the password with slat, the password generation Policy
//  @param plain saltPassword string
//  @param salt string
//  @return string
func saltPassword(plain string, salt string) string {
	return sha256x.Sha256(md5x.Md5x(salt + plain + salt))
}

// randomSalt generate random salt
// @return string
func randomSalt() string {
	return core.SimpleRandomUUID()
}
