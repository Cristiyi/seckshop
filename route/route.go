/**
 * @Author: cristi
 * @Description:
 * @File:  route.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package route

import (
	"github.com/kataras/iris/mvc"
	"seckshop/controllers/backend"
	"github.com/kataras/iris"
	"seckshop/controllers/fronted"

	//"github.com/kataras/iris/mvc"
	"net/http"
)
func InitRouter(app *iris.Application) {
	//app.Use(CrossAccess)
	bathPath := "/api/v1"
	mvc.New(app.Party(bathPath+"/admin/product")).Handle(backend.NewProductController())
	mvc.New(app.Party(bathPath+"/fronted/user")).Handle(fronted.NewUserController())
}
//bathPath := "/api/v1"

//func InitRouter(app *iris.Application) {
//	app.Use(CrossAccess)
//	bathPath := "/api/v1/admin"
//	ProductRouter(app, bathPath)
//}
//
//func ProductRouter(app *iris.Application, bathPath string) {
//	product := app.Party(bathPath + "/product")
//	mvc.New(product).Handle(backend.NewProductController())
//}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
