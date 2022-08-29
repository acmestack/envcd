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
	"errors"
	"fmt"
	"time"

	"github.com/acmestack/envcd/internal/core/storage/dao"
	"github.com/acmestack/envcd/internal/pkg/constant"
	"github.com/acmestack/envcd/internal/pkg/entity"
	"github.com/acmestack/envcd/internal/pkg/result"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-gonic/gin"
)

type scopeSpaceVO struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Note      string `json:"note"`
	State     string `json:"state"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Editable  bool   `json:"editable"`
}

type scopeSpaceDTO struct {
	ScopeSpaceName string `json:"scopeSpaceName"`
	Note           string `json:"note"`
	State          string `json:"state"`
}

func scopeSpaceConverter(scopeSpace entity.ScopeSpace, editable bool) scopeSpaceVO {
	return scopeSpaceVO{
		Id:        scopeSpace.Id,
		Name:      scopeSpace.Name,
		Note:      scopeSpace.Note,
		State:     scopeSpace.State,
		CreatedAt: scopeSpace.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: scopeSpace.UpdatedAt.Format("2006-01-02 15:04:05"),
		Editable:  editable,
	}
}

// scopeSpace get scope space by id
//  @receiver openapi openapi
//  @param ginCtx gin context
func (openapi *Openapi) scopeSpace(ginCtx *gin.Context) {
	openapi.execute(ginCtx, nil, func() *result.EnvcdResult {
		scopeSpaceId := stringsx.ToInt(ginCtx.Param("scopeSpaceId"))
		scopeSpace := entity.ScopeSpace{Id: scopeSpaceId}
		daoAction := dao.New(openapi.storage)
		scopeSpaceQueryRet, err := daoAction.SelectScopeSpace(scopeSpace)
		if err != nil {
			return result.InternalFailure(err)
		}
		// TODO can refactor
		if len(scopeSpaceQueryRet) == 0 {
			return result.InternalFailure(errors.New("no data"))
		}
		// TODO can refactor
		// query dictionary
		count, dictCountErr := daoAction.SelectDictionaryCount(entity.Dictionary{ScopeSpaceId: scopeSpaceId})
		if dictCountErr != nil {
			return result.InternalFailure(dictCountErr)
		}
		// no dictionary => scopeSpaceVO name can edit
		return result.Success(scopeSpaceConverter(scopeSpaceQueryRet[0], count == 0))
	})
}

func (openapi *Openapi) createScopeSpace(ginCtx *gin.Context) {
	openapi.execute(ginCtx, nil, func() *result.EnvcdResult {
		// query user have same scopeSpace for one person.
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil
	})
}

func (openapi *Openapi) updateScopeSpace(ginCtx *gin.Context) {
	openapi.execute(ginCtx, nil, func() *result.EnvcdResult {
		updateScopeSpace := &scopeSpaceDTO{}
		if err := ginCtx.ShouldBindJSON(updateScopeSpace); err != nil {
			fmt.Printf("Bind error, %v\n", err)
			return result.InternalFailure(err)
		}
		scopeSpaceId := stringsx.ToInt(ginCtx.Param("scopeSpaceId"))
		daoAction := dao.New(openapi.storage)
		scopeSpaceQueryRet, queryErr := daoAction.SelectScopeSpace(entity.ScopeSpace{Id: scopeSpaceId})
		if queryErr != nil {
			return result.InternalFailure(queryErr)
		}
		if len(scopeSpaceQueryRet) == 0 {
			return result.InternalFailure(errors.New("no scopespace error"))
		}
		defaultScopeSpace := scopeSpaceQueryRet[0]
		// name change, no dictionaries, just update mysql
		if defaultScopeSpace.Name != updateScopeSpace.ScopeSpaceName {
			updateRet, updateErr := daoAction.UpdateScopeSpace(entity.ScopeSpace{
				Id:        scopeSpaceId,
				Name:      updateScopeSpace.ScopeSpaceName,
				Note:      updateScopeSpace.Note,
				State:     updateScopeSpace.State,
				UpdatedAt: time.Now(),
			})
			if updateErr != nil {
				return result.InternalFailure(updateErr)
			}
			return result.Success(updateRet)
		}
		// name can't change, build no name scope space, judge state and update dictionary
		// scope space state change, following to do
		// 1.update scope space
		// 2.update all dictionaries
		if defaultScopeSpace.State == updateScopeSpace.State && defaultScopeSpace.Note == updateScopeSpace.Note {
			return result.Success(nil)
		}
		if defaultScopeSpace.State == updateScopeSpace.State && defaultScopeSpace.Note != updateScopeSpace.Note {
			updateRet, updateErr := daoAction.UpdateScopeSpace(entity.ScopeSpace{Note: updateScopeSpace.Note, UpdatedAt: time.Now()})
			if updateErr != nil {
				return result.InternalFailure(updateErr)
			}
			return result.Success(updateRet)
		}
		openapi.updateScopeSpaceState(defaultScopeSpace, updateScopeSpace.State)
		// if defaultScopeSpace.State != updateScopeSpace.State && defaultScopeSpace.Note == updateScopeSpace.Note
		// defaultScopeSpace.State != updateScopeSpace.State && defaultScopeSpace.Note != updateScopeSpace.Note
		// TODO update scope space state and update dictionary state and note need update
		return result.Success(nil)
	})
}

func (openapi *Openapi) removeScopeSpace(ginCtx *gin.Context) {
	openapi.execute(ginCtx, nil, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// TODO remove scopeSpace
		// 1.query dictionary by scopeSpaceId
		// 2.if no dictionary, then change scopeSpace state to deleted
		// 3.if there are some dictionaries, update scopeSpace to deleted and
		// change batch change dictionaries to deleted state and remove etcd data
		return nil
	})
}

func (openapi *Openapi) scopeSpaces(ginCtx *gin.Context) {
	openapi.execute(ginCtx, nil, func() *result.EnvcdResult {
		fmt.Println("hello world")
		// create config
		// ConfigDao.save();
		// go LogDao.save()
		// openapi.exchange.Put("key", "value")
		return nil
	})
}

func (openapi *Openapi) updateScopeSpaceState(defaultScopeSpace entity.ScopeSpace, newState string, note ...string) *result.EnvcdResult {
	if stringsx.Empty(newState) {
		return result.InternalFailure(errors.New("current state is nil"))
	}
	// defaultState must not equal new state
	// query exist dictionary
	daoAction := dao.New(openapi.storage)
	dictionary, queryDictErr := daoAction.SelectDictionary(entity.Dictionary{ScopeSpaceId: defaultScopeSpace.Id}, nil)
	if queryDictErr != nil {
		return result.InternalFailure(queryDictErr)
	}
	// no dictionary
	if len(dictionary) == 0 {
		var scopeSpace int64
		var updateErr error
		if defaultScopeSpace.Note == note[0] {
			scopeSpace, updateErr = daoAction.UpdateScopeSpace(entity.ScopeSpace{Id: defaultScopeSpace.Id, State: newState, UpdatedAt: time.Now()})
		} else {
			scopeSpace, updateErr = daoAction.UpdateScopeSpace(entity.ScopeSpace{Id: defaultScopeSpace.Id, Note: note[0], State: newState, UpdatedAt: time.Now()})
		}
		if updateErr != nil {
			return result.InternalFailure(updateErr)
		}
		return result.Success(scopeSpace)
	}

	switch newState {
	case constant.EnabledState:
		// if new state is enabled,change scopeSpace state and change dictionary state, end generate etcd path and put
		break
	case constant.DisabledState:
	case constant.DeletedState:
		break
	}
	return result.Failure0(result.ErrorNotExistState)
}
