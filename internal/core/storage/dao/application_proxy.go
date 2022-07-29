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
	modelV := entity.Application{}
	gobatis.RegisterModel(&modelV)
}

func SelectApplication(sess *gobatis.Session, model entity.Application) ([]entity.Application, error) {
	var dataList []entity.Application
	err := sess.Select("dao.selectApplication").Param(model).Result(&dataList)
	return dataList, err
}

func SelectApplicationCount(sess *gobatis.Session, model entity.Application) (int64, error) {
	var ret int64
	err := sess.Select("dao.selectApplicationCount").Param(model).Result(&ret)
	return ret, err
}

func InsertApplication(sess *gobatis.Session, model entity.Application) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertApplication").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func InsertBatchApplication(sess *gobatis.Session, models []entity.Application) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertBatchApplication").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func UpdateApplication(sess *gobatis.Session, model entity.Application) (int64, error) {
	var ret int64
	err := sess.Update("dao.updateApplication").Param(model).Result(&ret)
	return ret, err
}

func DeleteApplication(sess *gobatis.Session, model entity.Application) (int64, error) {
	var ret int64
	err := sess.Delete("dao.deleteApplication").Param(model).Result(&ret)
	return ret, err
}
