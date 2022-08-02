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

package data

import (
	"encoding/json"
	"log"
)

const (
	StringFormat     = "string"
	JsonFormat       = "json"
	YamlFormat       = "yaml"
	PropertiesFormat = "properties"
)

// EnvcdData for dictionary value
type EnvcdData struct {
	Format string      `json:"format"`
	Data   interface{} `json:"data"`
}

// String create string format envcd data
//  @param data of string format
func String(data interface{}) EnvcdData {
	return EnvcdData{
		Format: StringFormat,
		Data:   data,
	}
}

// Json create json format envcd data
//  @param data of json format
func Json(data interface{}) EnvcdData {
	return EnvcdData{
		Format: JsonFormat,
		Data:   data,
	}
}

// Yaml create yaml format envcd data
//  @param data of yaml format
func Yaml(data interface{}) EnvcdData {
	return EnvcdData{
		Format: YamlFormat,
		Data:   data,
	}
}

// Properties create properties format envcd data
//  @param data of properties format
func Properties(data interface{}) EnvcdData {
	return EnvcdData{
		Format: PropertiesFormat,
		Data:   data,
	}
}

// ToJson envcd data convert to json string
func ToJson(envcdData EnvcdData) string {
	marshal, err := json.Marshal(envcdData)
	if err != nil {
		log.Fatalln(err)
	}
	return string(marshal)
}

// ToEnvcdData json string convert to EnvcdData
func ToEnvcdData(jsonString string) EnvcdData {
	data := EnvcdData{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
