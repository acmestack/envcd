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

package storage

import (
	"embed"
	"fmt"
	"log"

	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis/datasource"
	"github.com/acmestack/gobatis/factory"
	"github.com/acmestack/godkits/gox/errorsx"
	"github.com/acmestack/pagehelper"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	storage        *config.Storage
	sessionManager *gobatis.SessionManager
}

func Start(mysql *config.Storage) *Storage {
	// load sqlmap
	loadSqlMap()
	// create SessionManager
	dbFactory := initFactory(mysql)
	pageHelpFactory := pagehelper.New(dbFactory)
	return &Storage{storage: mysql, sessionManager: gobatis.NewSessionManager(pageHelpFactory)}
}

// NewSession new session
//  @return *gobatis.Session
func (storage *Storage) NewSession() *gobatis.Session {
	// todo storage check
	if storage == nil {
		log.Fatalln(errorsx.Err("IIllegal state for storage"))
	}
	return storage.sessionManager.NewSession()
}

// InitDB init sql session manager
//  @param mysql config
//  @return *gobatis.SessionManager sessionManager
func initFactory(mysql *config.Storage) factory.Factory {
	return gobatis.NewFactory(
		gobatis.SetMaxConn(100),
		gobatis.SetMaxIdleConn(50),
		gobatis.SetDataSource(&datasource.MysqlDataSource{
			Host:     mysql.ConnMetadata.Hostname,
			Port:     mysql.ConnMetadata.Port,
			DBName:   mysql.Database,
			Username: mysql.ConnMetadata.UserName,
			Password: mysql.ConnMetadata.Password,
			Charset:  "utf8",
		}))
}

//go:embed xml/*.xml
var locationFS embed.FS

// loadSqlMap load sql map from directory
func loadSqlMap() {
	locationFiles, err := locationFS.ReadDir("xml")
	if err != nil {
		fmt.Println("read xml dir is error:", err.Error())
		return
	}
	for _, location := range locationFiles {
		file, err := locationFS.ReadFile("xml/" + location.Name())
		if err != nil {
			fmt.Println("reade file is error:", err.Error())
			continue
		}
		err = gobatis.RegisterMapperData(file)
		if err != nil {
			fmt.Println("parse mappers is error:", err.Error())
		}
	}
}
