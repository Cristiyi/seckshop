/**
 * @Author: cristi
 * @Description:jwt
 * @File:  jwt.go
 * @Version: 1.0.0
 * @Date: 2019/12/11 0010 14:08
 */

package middleware

import (
	//"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"seckshop/models"
	"time"
	"github.com/kataras/iris/v12"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
var JwtHandler = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,

})



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
//func CheckToken(ctx iris.Context) (isRight bool, err error){
//
//	tokenObj, _ := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
//		if err != nil {
//			isRight = false
//			return
//		}
//		return "secret", nil
//	})
//
//	//校验错误（基本）
//	err = tokenObj.Claims.Valid()
//	if err != nil {
//		return false, err
//	}
//
//	finToken := tokenObj.Claims.(jwt.MapClaims)
//	//校验下token是否过期
//	succ := finToken.VerifyExpiresAt(time.Now().Unix(),true)
//	fmt.Println("succ",succ)
//	fmt.Println(finToken)
//	return true, nil
//
//}

func ParseToken(ctx iris.Context) {

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "need token"})
		return
	}
	tokenObj, _ := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if err != nil {
			ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "token error"})
			return
		}
		return "secret", nil
	})

	////校验错误（基本）
	//validErr := tokenObj.Claims.Valid()
	//if validErr != nil {
	//	ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "token error"})
	//	return
	//}

	finToken := tokenObj.Claims.(jwt.MapClaims)
	//校验下token是否过期
	isSucc := finToken.VerifyExpiresAt(time.Now().Unix(),true)
	if isSucc {
		ctx.Next()
	}

	ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "token error"})
	return

}
