/**
 * @Author: cristi
 * @Description:
 * @File:  user_controller.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package fronted

import (
	"github.com/kataras/iris"
	"seckshop/models"
	"seckshop/services"
	"seckshop/utils"
)

type UserController struct {
	Ctx iris.Context
	Service services.UserService
}

func NewUserController() *UserController {
	return &UserController{Service:services.NewUserService()}
}

func (g *UserController) PostInsert() (result models.Result) {

	r := g.Ctx.Request()

	tel := r.PostFormValue("tel")
	if tel == "" {
		result.Msg = "手机号不能为空"
		result.Code = 500
		return
	}
	checkTel := utils.VerifyMobileFormat(tel)
	if checkTel == false {
		result.Msg = "请传入正确手机号"
		result.Code = 500
		return
	}
	userName := r.PostFormValue("username")
	if userName == "" {
		result.Msg = "用户名不能为空"
		result.Code = 500
		return
	}
	nickName := r.PostFormValue("nickname")
	if nickName == "" {
		result.Msg = "昵称不能为空"
		result.Code = 500
		return
	}
	password := r.PostFormValue("password")
	if password == "" {
		result.Msg = "密码不能为空"
		result.Code = 500
		return
	}
	m := make(map[string]interface{})
	m["username"] = userName
	m["nickname"] = nickName
	m["password"] = password
	m["tel"] = tel
	return g.Service.Insert(m)

}

func (u *UserController) PostLogin() (result models.Result) {
	r := u.Ctx.Request()
	tel := r.PostFormValue("tel")
	if tel == "" {
		result.Msg = "手机号不能为空"
		result.Code = 500
		return
	}
	password := r.PostFormValue("password")
	if password == "" {
		result.Msg = "密码不能为空"
		result.Code = 500
		return
	}
	return u.Service.CheckLogin(tel, password)
}
