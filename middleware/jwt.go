/**
 * @Author: cristi
 * @Description:jwt
 * @File:  jwt.go
 * @Version: 1.0.0
 * @Date: 2019/12/11 0010 14:08
 */

package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jailer "github.com/iris-contrib/middleware/jwt"
	"time"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func JwtHandler() *jailer.Middleware {
	return jailer.New(jailer.Config {
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

//创建token
func CreateToken(userId int64) (token string, err error) {

	tokenObj := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = userId
	tokenObj.Claims = claims
	token, err = tokenObj.SignedString([]byte("secret"))
	return

}

//验证token
func CheckToken(tokenString string, key string) (isRight bool) {
	tokenObj, _ := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if err != nil {
			isRight = false
			return
		}
		return "secret", nil
	})
	
	//校验错误（基本）
	err := tokenObj.Claims.Valid()
	if err != nil {
		return false
	}

	finToken := tokenObj.Claims.(jwt.MapClaims)
	//校验下token是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(),true)
	fmt.Println("succ",succ)
	fmt.Println(finToken)
	return true

}
