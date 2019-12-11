/**
 * @Author: cristi
 * @Description:jwt
 * @File:  jwt.go
 * @Version: 1.0.0
 * @Date: 2019/12/11 0010 14:08
 */

package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jailer "github.com/iris-contrib/middleware/jwt"
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
