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

package dao

import (
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/gobatis"
)

func init() {
	modelV := entity.Logging{}
	gobatis.RegisterModel(&modelV)
}

func SelectLogging(sess *gobatis.Session, model entity.Logging) ([]entity.Logging, error) {
	var dataList []entity.Logging
	err := sess.Select("dao.selectLogging").Param(model).Result(&dataList)
	return dataList, err
}

func SelectLoggingCount(sess *gobatis.Session, model entity.Logging) (int64, error) {
	var ret int64
	err := sess.Select("dao.selectLoggingCount").Param(model).Result(&ret)
	return ret, err
}

func InsertLogging(sess *gobatis.Session, model entity.Logging) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertLogging").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func InsertBatchLogging(sess *gobatis.Session, models []entity.Logging) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertBatchLogging").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func UpdateLogging(sess *gobatis.Session, model entity.Logging) (int64, error) {
	var ret int64
	err := sess.Update("dao.updateLogging").Param(model).Result(&ret)
	return ret, err
}

func DeleteLogging(sess *gobatis.Session, model entity.Logging) (int64, error) {
	var ret int64
	err := sess.Delete("dao.deleteLogging").Param(model).Result(&ret)
	return ret, err
}
