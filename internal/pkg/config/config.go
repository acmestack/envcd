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

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

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

// NewConfig new envcd config
//  @param configFile the config file
//  @return *Config current config instance
func NewConfig(configFile *string) *Config {
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	envcdConfig := &Config{}
	if e := yaml.Unmarshal(data, envcdConfig); e != nil {
		log.Fatalf("error %v", err)
	}
	return envcdConfig
}

// Information the envcd config information
//  @receiver cfg
func (cfg *Config) Information() {
	// todo log
	fmt.Println(cfg.Exchanger)
}
