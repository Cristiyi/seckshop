/**
 * @Author: cristi
 * @Description:
 * @File:  user_service.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package services

import (
	"github.com/spf13/cast"
	"seckshop/models"
	"seckshop/repo"
)

type UserService interface {
	Insert(m map[string]interface{}) (result models.Result)
	CheckLogin(tel string, password string) (userId int64, result models.Result)
}

type userService struct {
	Repository		repo.UserRepo
}

func NewUserService() UserService {
	return &userService{Repository: repo.NewUserRepo()}
}


func (s userService) Insert(m map[string]interface{}) (result models.Result) {

	isReg, checkErr := s.Repository.CheckTel(cast.ToString(m["tel"]))
	if checkErr != nil {
		result.Data = nil
		result.Code = 500
		result.Msg = checkErr.Error()
		return
	}
	if isReg == true {
		result.Data = nil
		result.Code = 500
		result.Msg = "添加失败，该手机号已注册"
		return
	}
	userId, insertErr := s.Repository.Insert(m)
	if insertErr != nil {
		result.Data = nil
		result.Code = 500
		result.Msg = insertErr.Error()
		return
	}
	result.Data = userId
	result.Code = 200
	result.Msg = "添加成功"
	return

}

func (s userService) CheckLogin(tel string, password string) (userId int64, result models.Result) {
	isRight, data, msg := s.Repository.CheckLogin(tel, password)
	if !isRight {
		userId = 0
		result.Data = nil
		result.Code = 500
		result.Msg =  msg
		return
	}
	maps := make(map[string]interface{}, 1)
	maps["user"] = data
	result.Msg = "登陆成功"
	result.Code = 200
	result.Data = maps
	userId = data.ID
	return

}
