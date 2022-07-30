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

package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis/datasource"
	"github.com/acmestack/gobatis/factory"
	"github.com/acmestack/godkits/gox/errorsx"
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
	db := initDB(mysql)
	return &Storage{storage: mysql, sessionManager: gobatis.NewSessionManager(db)}
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
func initDB(mysql *config.Storage) factory.Factory {
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

// loadSqlMap load sql map from directory
func loadSqlMap() {
	envcdPath, err := os.Getwd()
	if err != nil {
		fmt.Println("get envcd path is error:", err.Error())
	}
	err = gobatis.ScanMapperFile(envcdPath + "/internal/core/storage/xml")
	if err != nil {
		fmt.Println("parse mappers is error:", err.Error())
	}
}
