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

package configuration

import "github.com/acmestack/envcd/pkg/entity"

type Configuration struct {
	user  *entity.User
	space *entity.ScopeSpace
	data  *entity.Envcdata
}

// NewConfiguration create new Configuration by user, space, data.
//  @param user the Configuration owner
//  @param space the Configuration space
//  @param data the Configuration data
//  @return *Configuration
func NewConfiguration(user *entity.User, space *entity.ScopeSpace, data *entity.Envcdata) *Configuration {
	return &Configuration{
		user:  user,
		space: space,
		data:  data,
	}
}

// Identity figure the configuration's identity.
// identity = user + space todo
//  @receiver configuration current config
//  @return string identity the configuration's empty able identity
func (configuration *Configuration) Identity() string {
	if configuration == nil {
		return ""
	}
	// todo
	return ""
}

// Data get configuration's data.
//  @receiver configuration current config
//  @return *entity.Envcdata the configuration's nullable data
func (configuration *Configuration) Data() *entity.Envcdata {
	if configuration == nil {
		return nil
	}
	return configuration.data
}
