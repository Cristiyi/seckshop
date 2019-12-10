/**
 * @Author: cristi
 * @Description:
 * @File:  user_repo.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package repo

import (
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"seckshop/models"
)

type UserRepo interface {
	Insert(m map[string]interface{}) (userId int64, err error)
	CheckTel(tel string) (isReg bool, err error)
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}


type userRepo struct {

}

func(u userRepo) Insert(m map[string]interface{}) (userId int64, err error) {

	user := new(models.User)
	user.Tel = cast.ToString(m["tel"])
	user.NickName = cast.ToString(m["nickname"])
	user.UserName = cast.ToString(m["username"])
	pwdByte, pwdError := GeneratePassword(cast.ToString(m["password"]))
	if pwdError != nil {
		panic("gen pwd error")
	}
	user.HashPassword = string(pwdByte)
	//user.CreatedAt = time.Now()
	_, err = engine.Insert(user)
	if err != nil {
		panic("insert error")
	}
	userId = user.ID
	return

}

//检查用户手机号是否注册
func(u userRepo) CheckTel(tel string) (isReg bool, err error) {
	has, err := engine.Where("tel = ?", tel).Get(&models.User{})
	if err != nil {
		panic("error")
	}
	isReg = has
	return
}

//生成password
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}