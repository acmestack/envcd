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

package openapi

import (
	"fmt"
	"net/http"
	"time"

	authjwt "github.com/acmestack/envcd/internal/core/jwt"
	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/pkg/entity/result"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/acmestack/godkits/log"
	"github.com/gin-gonic/gin"
)

// loginParam Login
type loginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// userParam Create User Param
type userParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Identity int    `json:"identity"`
	State    bool   `json:"state"`
}

func (openapi *Openapi) login(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		param := loginParam{}
		if err := ginCtx.ShouldBindJSON(&param); err != nil {
			log.Error("Bind error, %v", err)
			return result.InternalServerErrorFailure("Illegal params !")
		}

		users, err := dao.New(openapi.storage).SelectUser(entity.User{
			Name: param.Username,
		})
		if err != nil {
			log.Error("Query User error: %v", err)
			// todo error code : result.Failure0(code, message, httpStatusCode)
			return result.Failure("System Error!", http.StatusBadRequest)
		}

		if len(users) == 0 {
			// todo error code : result.Failure0(code, message, httpStatusCode)
			log.Error("User does not exist : %v", param)
			return result.Failure("User does not exist!", http.StatusOK)
		}
		user := users[0]
		if saltPassword(param.Password, user.Salt) != user.Password {
			// todo error code : result.Failure0(code, message, httpStatusCode)
			return result.Failure("password error!", http.StatusOK)
		}
		token := authjwt.NewJWTToken(authjwt.AuthClaims{
			RegisteredClaims: &jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Second)),
			},
			UserId:   user.Id,
			UserName: user.Name,
		})
		return result.Success(map[string]interface{}{
			// todo const var
			"userId": user.Id,
			"token":  token,
		})
	})
}

func (openapi *Openapi) logout(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// UserDao.save(),
		// LogDao.save()
		return nil
	})
}

func (openapi *Openapi) user(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		param := userParam{}
		if er := ginCtx.ShouldBindJSON(&param); er != nil {
			log.Error("Bind error, %v", er)
			return result.InternalServerErrorFailure("Illegal params !")
		}
		daoApi := dao.New(openapi.storage)
		// check if the user already exists in the database
		users, er := daoApi.SelectUser(entity.User{
			Name: param.Name,
		})
		if er != nil {
			log.Error("Query User error: %v", er)
			return result.InternalServerErrorFailure("System Error!")
		}
		if len(users) > 0 {
			log.Error("User Has exists: %v", users)
			return result.InternalServerErrorFailure("User Has Exists!")
		}
		// generate database password by salt
		salt := randomSalt()
		password := saltPassword(param.Password, salt)
		state := 1
		if !param.State {
			state = 2
		}
		user := entity.User{
			Name:      param.Name,
			Password:  password,
			Salt:      salt,
			Identity:  param.Identity,
			State:     state,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// save user
		if _, _, err := daoApi.InsertUser(user); err != nil {
			log.Error("insert error=%v", err)
			return result.InternalServerErrorFailure("Save User Error!")
		}
		// fixme update success message or response token and id ?
		return result.Success("ok")
	})
}

func (openapi *Openapi) userById(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		id := stringsx.ToInt(ginCtx.Param("id"))
		user := entity.User{Id: id}
		// todo user detail
		dao.New(openapi.storage).SelectUser(user)
		return nil
	})
}

func (openapi *Openapi) removeUser(ginCtx *gin.Context) {
	openapi.response(ginCtx, func() *result.EnvcdResult {
		fmt.Println("hello world")
		return nil
	})
}
