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
	modelV := entity.Permission{}
	gobatis.RegisterModel(&modelV)
}

func SelectPermission(sess *gobatis.Session, model entity.Permission) ([]entity.Permission, error) {
	var dataList []entity.Permission
	err := sess.Select("dao.selectPermission").Param(model).Result(&dataList)
	return dataList, err
}

func SelectPermissionCount(sess *gobatis.Session, model entity.Permission) (int64, error) {
	var ret int64
	err := sess.Select("dao.selectPermissionCount").Param(model).Result(&ret)
	return ret, err
}

func InsertPermission(sess *gobatis.Session, model entity.Permission) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertPermission").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func InsertBatchPermission(sess *gobatis.Session, models []entity.Permission) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertBatchPermission").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func UpdatePermission(sess *gobatis.Session, model entity.Permission) (int64, error) {
	var ret int64
	err := sess.Update("dao.updatePermission").Param(model).Result(&ret)
	return ret, err
}

func DeletePermission(sess *gobatis.Session, model entity.Permission) (int64, error) {
	var ret int64
	err := sess.Delete("dao.deletePermission").Param(model).Result(&ret)
	return ret, err
}
