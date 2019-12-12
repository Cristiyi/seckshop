/**
 * @Author: cristi
 * @Description:
 * @File:  user_repo.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package repo

import (
	"fmt"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"seckshop/middleware"
	"seckshop/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserRepo interface {
	Insert(m map[string]interface{}) (userId int64, err error)
	CheckTel(tel string) (isReg bool, err error)
	CheckLogin(tel string, password string) (has bool, token string, msg string)
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
//登录逻辑
func(u userRepo) CheckLogin(tel string, password string) (has bool, token string, msg string) {

	user := new(models.User)
	has, err := engine.Where("tel=?", tel).Get(user)
	if err != nil {
		panic("select user error")
	}
	if !has {
		return false, "", "未找到用户"
	}

	getPassword := ValidatePassword(password, user.HashPassword);
	if !getPassword {
		return false, "", "密码错误，请重新输入"
	}

	tokenObj := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	tokenObj.Claims = claims
	token, err = middleware.CreateToken(user.ID)

	if err != nil {
		panic("token error")
	}
	msg = "登陆成功"
	user.Token = token
	affected, err := engine.Id(user.ID).Cols("token").Update(user)
	if err != nil {
		fmt.Println(err.Error())
		panic("token error")
	}
	middleware.CheckToken(token, "secret")
	if affected == 0 {
		return false, "", "登录失败"
	}
	return
}

//生成password
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//验证password
func ValidatePassword(password string, hashedPassword string) (isRight bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}