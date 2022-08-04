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

package dao

import (
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/gobatis"
)

func init() {
	gobatis.RegisterModel(&entity.Dictionary{})
}

func (dao *Dao) SelectDictionary(model entity.Dictionary) ([]entity.Dictionary, error) {
	var dataList []entity.Dictionary
	err := dao.storage.NewSession().Select("dao.selectDictionary").Param(model).Result(&dataList)
	return dataList, err
}

func (dao *Dao) SelectDictionaryCount(model entity.Dictionary) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Select("dao.selectDictionaryCount").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) InsertDictionary(model entity.Dictionary) (int64, int64, error) {
	var ret int64
	runner := dao.storage.NewSession().Insert("dao.insertDictionary").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func (dao *Dao) InsertBatchDictionary(models []entity.Dictionary) (int64, int64, error) {
	var ret int64
	runner := dao.storage.NewSession().Insert("dao.insertBatchDictionary").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func (dao *Dao) UpdateDictionary(model entity.Dictionary) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Update("dao.updateDictionary").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) DeleteDictionary(model entity.Dictionary) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Delete("dao.deleteDictionary").Param(model).Result(&ret)
	return ret, err
}
