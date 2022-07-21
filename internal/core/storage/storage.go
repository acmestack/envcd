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
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis/datasource"
)

type Storage struct {
	storage *config.Storage
}

func Start(mysql *config.Storage) *Storage {
	return &Storage{storage: mysql}
}

// InitDB init sql session manager
//  @param mysql config
//  @return *gobatis.SessionManager sessionManager
func InitDB(mysql *config.Storage) *gobatis.SessionManager {
	fac := gobatis.NewFactory(
		gobatis.SetMaxConn(100),
		gobatis.SetMaxIdleConn(50),
		gobatis.SetDataSource(&datasource.MysqlDataSource{
			Host:     mysql.ConnMetadata.Host,
			Port:     mysql.ConnMetadata.Port,
			DBName:   mysql.Database,
			Username: mysql.ConnMetadata.UserName,
			Password: mysql.ConnMetadata.Password,
			Charset:  "utf8",
		}))
	return gobatis.NewSessionManager(fac)
}
