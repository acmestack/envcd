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
	"github.com/acmestack/godkits/log"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	exchangerType = "Exchanger"
	mysqlType     = "Storage"
)

// Exchanger the Exchanger config
type Exchanger struct {
	// Exchanger with standard Url: etcd://user:123@localhost:123
	// the schema is the kind of the center
	Url          string `yaml:"url"`
	ConnMetadata *ConnMetadata
}

// Storage the Storage config
type Storage struct {
	// Url with standard Url: MySQL://user:123@localhost:123
	Url          string `yaml:"url"`
	Database     string `yaml:"database"`
	ConnMetadata *ConnMetadata
}

// Server the Server config
type Server struct {
	RunMode      string `yaml:"run-mode"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read-timeout"`
	WriteTimeout int    `yaml:"write-timeout"`
}

// Config the envcd config
type Config struct {
	Exchanger *Exchanger `yaml:"exchanger"`
	Storage   *Storage   `yaml:"storage"`
	Server    *Server    `yaml:"server"`
}

// NewConfig new envcd config
//  @param configFile the config file
//  @return *Config current config instance
func NewConfig(configFile *string) *Config {
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Error("error %s", err)
	}
	envcdConfig := &Config{}
	if e := yaml.Unmarshal(data, envcdConfig); e != nil {
		log.Error("error %s", err)
	}
	return envcdConfig
}

// StartInformation the envcd config information
//  @receiver cfg
func (cfg *Config) StartInformation() {
	cfg.Exchanger.ConnMetadata = parser(cfg.Exchanger.Url)
	cfg.Exchanger.ConnMetadata.information(exchangerType)
	cfg.Storage.ConnMetadata = parser(cfg.Storage.Url)
	cfg.Storage.ConnMetadata.information(mysqlType)
}
