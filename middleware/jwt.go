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
	"seckshop/models"
	"strings"
	"time"
	"github.com/kataras/iris/v12"
)

const JWTKEY = "9e95bf56a1e3dcb44a34ae7fc9034091"

//jwt Claim struct
type Claims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	UserName string `json:"username"`
}

//创建token
func CreateToken(user models.User) (token string, err error) {
	claims := &Claims{
		UserId:   user.ID,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Hour*time.Duration(1)).Unix(),
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenObj.SignedString([]byte(JWTKEY))
	return
}

//验证token
func ParseToken(ctx iris.Context) {

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "need token"})
		return
	}
	tokenString = strings.Split(tokenString, "Bearer ")[1]
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKEY), nil
	})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: err.Error()})
		return
	}
	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "That's not even a token"})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "token is expired"})
			} else {
				ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "token is invalid"})
			}
			return
		}
		ctx.JSON(&models.Result{Code: 500, Data: nil, Msg: "token error2"})
		return
	}
	ctx.Next()
	return

}
