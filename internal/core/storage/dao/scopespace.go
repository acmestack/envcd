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
	gobatis.RegisterModel(&entity.ScopeSpace{})
}

func (dao *Dao) SelectScopeSpace(model entity.ScopeSpace) ([]entity.ScopeSpace, error) {
	var dataList []entity.ScopeSpace
	err := dao.storage.NewSession().Select("dao.selectScopeSpace").Param(model).Result(&dataList)
	return dataList, err
}

func (dao *Dao) SelectScopeSpaceCount(model entity.ScopeSpace) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Select("dao.selectScopeSpaceCount").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) InsertScopeSpace(model entity.ScopeSpace) (int64, int64, error) {
	var ret int64
	runner := dao.storage.NewSession().Insert("dao.insertScopeSpace").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func (dao *Dao) InsertBatchScopeSpace(models []entity.ScopeSpace) (int64, int64, error) {
	var ret int64
	runner := dao.storage.NewSession().Insert("dao.insertBatchScopeSpace").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func (dao *Dao) UpdateScopeSpace(model entity.ScopeSpace) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Update("dao.updateScopeSpace").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) UpdateScopeSpaceBatch(model []entity.ScopeSpace) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Update("dao.updateScopeSpaceBatch").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) DeleteScopeSpace(model entity.ScopeSpace) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Delete("dao.deleteScopeSpace").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) DeleteScopeSpaceBatch(model []entity.ScopeSpace) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Delete("dao.deleteScopeSpaceBatch").Param(model).Result(&ret)
	return ret, err
}
