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
	authjwt "github.com/acmestack/envcd/internal/core/jwt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"

	"github.com/acmestack/envcd/internal/core/plugin"
	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/pkg/entity/data"
	"github.com/acmestack/godkits/gox/errorsx"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/acmestack/godkits/log"
	"github.com/gin-gonic/gin"
)

// loginParam Login
type loginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (openapi *Openapi) login(ctx *gin.Context) {
	// receive params from request
	param := loginParam{}
	if er := ctx.ShouldBindJSON(&param); er != nil {
		log.Error("Bind error, %v", er)
		ctx.JSON(http.StatusInternalServerError, data.Failure("Illegal params !").Data)
		return
	}
	daoApi := dao.New(openapi.storage)
	users, er := daoApi.SelectUser(entity.User{
		Name: param.Username,
	})
	if er != nil {
		log.Error("Query User error: %v", er)
		ctx.JSON(http.StatusBadRequest, data.Failure("System Error!").Data)
		return
	}
	if len(users) == 0 {
		log.Error("User does not exist : %v", param)
		ctx.JSON(http.StatusBadRequest, data.Failure("User does not exist!").Data)
		return
	}
	user := users[0]
	if saltPassword(param.Password, user.Salt) != user.Password {
		ctx.JSON(http.StatusBadRequest, data.Failure("password error!").Data)
		return
	}
	token := authjwt.NewJWTToken(authjwt.AuthClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Second)),
		},
		UserId:   user.Id,
		UserName: user.Name,
	})
	ctx.JSON(200, data.Success(map[string]interface{}{
		"userId": user.Id,
		"token":  token,
	}).Data)
}

func (openapi *Openapi) logout(ctx *gin.Context) {
	c, _ := buildContext(ctx)
	c.Action = func() (*data.EnvcdResult, error) {
		fmt.Println("hello world")
		// UserDao.save(),
		// LogDao.save()
		return nil, errorsx.Err("test error")
	}
	if ret, err := plugin.NewChain(openapi.executors).Execute(c); err != nil {
		fmt.Printf("ret = %v, error = %v", ret, err)
	}
	ctx.JSON(200, data.Success("hello world").Data)
}

// userParam Create User Param
type userParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Identity int    `json:"identity"`
	State    bool   `json:"state"`
}

func (openapi *Openapi) user(ctx *gin.Context) {
	c := &context.Context{Action: func() (*data.EnvcdResult, error) {
		fmt.Println("hello world")
		return nil, errorsx.Err("test error")
	}}
	if ret, err := plugin.NewChain(openapi.executors).Execute(c); err != nil {
		fmt.Printf("ret = %v, error = %v \n", ret, err)
	}
	// receive params from request
	param := userParam{}
	if er := ctx.ShouldBindJSON(&param); er != nil {
		log.Error("Bind error, %v", er)
		ctx.JSON(http.StatusInternalServerError, data.Failure("Illegal params !").Data)
		return
	}
	daoApi := dao.New(openapi.storage)
	// check if the user already exists in the database
	users, er := daoApi.SelectUser(entity.User{
		Name: param.Name,
	})
	if er != nil {
		log.Error("Query User error: %v", er)
		ctx.JSON(http.StatusInternalServerError, data.Failure("System Error!").Data)
		return
	}
	if len(users) > 0 {
		log.Error("User Has exists: %v", users)
		ctx.JSON(http.StatusInternalServerError, data.Failure("User Has Exists!").Data)
		return
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
		ctx.JSON(http.StatusInternalServerError, data.Failure("Save User Error!").Data)
		return
	}
	ctx.JSON(http.StatusOK, data.Success(nil).Data)
}

func (openapi *Openapi) userById(ctx *gin.Context) {
	c := &context.Context{Action: func() (*data.EnvcdResult, error) {
		fmt.Println("hello world")
		return nil, errorsx.Err("test error")
	}}
	id := stringsx.ToInt(ctx.Param("id"))
	user := entity.User{Id: id}
	dao.New(openapi.storage).SelectUser(user)
	if ret, err := plugin.NewChain(openapi.executors).Execute(c); err != nil {
		fmt.Printf("ret = %v, error = %v", ret, err)
	}
	ctx.JSON(200, data.Success("hello world").Data)
}

func (openapi *Openapi) removeUser(ctx *gin.Context) {
	c := &context.Context{Action: func() (*data.EnvcdResult, error) {
		fmt.Println("hello world")
		return nil, errorsx.Err("test error")
	}}
	if ret, err := plugin.NewChain(openapi.executors).Execute(c); err != nil {
		fmt.Printf("ret = %v, error = %v", ret, err)
	}
	ctx.JSON(200, data.Success("hello world").Data)
}
