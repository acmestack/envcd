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

package config

// mysql the MySQL config
type mysql struct {
	// Url with standard Url: mysql://user:123@localhost:123
	Url string `yaml:"url"`
}

// Config the envcd config
type Config struct {
	// Exchanger with standard Url: etcd://user:123@localhost:123
	// the schema is the kind of the center
	Exchanger string `yaml:"exchanger"`
	Mysql     *mysql `yaml:"mysql"`
}
