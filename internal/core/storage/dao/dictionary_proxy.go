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
	modelV := entity.Dictionary{}
	gobatis.RegisterModel(&modelV)
}

func SelectDictionary(sess *gobatis.Session, model entity.Dictionary) ([]entity.Dictionary, error) {
	var dataList []entity.Dictionary
	err := sess.Select("dao.selectDictionary").Param(model).Result(&dataList)
	return dataList, err
}

func SelectDictionaryCount(sess *gobatis.Session, model entity.Dictionary) (int64, error) {
	var ret int64
	err := sess.Select("dao.selectDictionaryCount").Param(model).Result(&ret)
	return ret, err
}

func InsertDictionary(sess *gobatis.Session, model entity.Dictionary) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertDictionary").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func InsertBatchDictionary(sess *gobatis.Session, models []entity.Dictionary) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertBatchDictionary").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func UpdateDictionary(sess *gobatis.Session, model entity.Dictionary) (int64, error) {
	var ret int64
	err := sess.Update("dao.updateDictionary").Param(model).Result(&ret)
	return ret, err
}

func DeleteDictionary(sess *gobatis.Session, model entity.Dictionary) (int64, error) {
	var ret int64
	err := sess.Delete("dao.deleteDictionary").Param(model).Result(&ret)
	return ret, err
}
