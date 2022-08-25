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
	gobatis.RegisterModel(&entity.Permission{})
}

func (dao *Dao) SelectPermission(model entity.Permission) ([]entity.Permission, error) {
	var dataList []entity.Permission
	err := dao.storage.NewSession().Select("dao.selectPermission").Param(model).Result(&dataList)
	return dataList, err
}

func (dao *Dao) SelectPermissionCount(model entity.Permission) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Select("dao.selectPermissionCount").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) InsertPermission(model entity.Permission) (int64, int64, error) {
	var ret int64
	runner := dao.storage.NewSession().Insert("dao.insertPermission").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func (dao *Dao) InsertBatchPermission(models []entity.Permission) (int64, int64, error) {
	var ret int64
	runner := dao.storage.NewSession().Insert("dao.insertBatchPermission").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func (dao *Dao) UpdatePermission(model entity.Permission) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Update("dao.updatePermission").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) UpdatePermissionBatch(model []entity.Permission) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Update("dao.updatePermissionBatch").Param(model).Result(&ret)
	return ret, err
}

func (dao *Dao) DeletePermission(model entity.Permission) (int64, error) {
	var ret int64
	err := dao.storage.NewSession().Delete("dao.deletePermission").Param(model).Result(&ret)
	return ret, err
}
