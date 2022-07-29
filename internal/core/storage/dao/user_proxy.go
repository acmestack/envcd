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
	modelV := entity.User{}
	gobatis.RegisterModel(&modelV)
}

func SelectUser(sess *gobatis.Session, model entity.User) ([]entity.User, error) {
	var dataList []entity.User
	err := sess.Select("dao.selectUser").Param(model).Result(&dataList)
	return dataList, err
}

func SelectUserCount(sess *gobatis.Session, model entity.User) (int64, error) {
	var ret int64
	err := sess.Select("dao.selectUserCount").Param(model).Result(&ret)
	return ret, err
}

func InsertUser(sess *gobatis.Session, model entity.User) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertUser").Param(model)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func InsertBatchUser(sess *gobatis.Session, models []entity.User) (int64, int64, error) {
	var ret int64
	runner := sess.Insert("dao.insertBatchUser").Param(models)
	err := runner.Result(&ret)
	id := runner.LastInsertId()
	return ret, id, err
}

func UpdateUser(sess *gobatis.Session, model entity.User) (int64, error) {
	var ret int64
	err := sess.Update("dao.updateUser").Param(model).Result(&ret)
	return ret, err
}

func DeleteUser(sess *gobatis.Session, model entity.User) (int64, error) {
	var ret int64
	err := sess.Delete("dao.deleteUser").Param(model).Result(&ret)
	return ret, err
}
