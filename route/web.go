package route

import (
	"github.com/kataras/iris/v12"
	"seckshop/controllers/backend"
	"seckshop/middleware"

	"github.com/kataras/iris/v12/core/router"
	"seckshop/controllers/fronted"
)


func Register(api *iris.Application) {


	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	//	AllowedHeaders:   []string{"*"},
	//	ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	//	AllowCredentials: true,
	//})


	v1 := api.Party("/v1").AllowMethods(iris.MethodOptions)
	{
		//front
		v1.PartyFunc("/front", func(front router.Party) {
			front.Post("/login", fronted.Login)
			front.PartyFunc("/product", func(product router.Party) {
				//product.Use(middleware.JwtHandler.Serve)
				product.Use(middleware.ParseToken)
				product.Post("/detail", fronted.Detail)
			})
			front.PartyFunc("/order", func(order router.Party) {
				//product.Use(middleware.JwtHandler.Serve)
				order.Use(middleware.ParseToken)
				order.Post("/seck", fronted.Seck)
			})
		})

		//admin
		v1.PartyFunc("/admin", func(admin router.Party) {
			admin.PartyFunc("/product", func(product router.Party) {
				product.Post("/list", backend.ProductList)
				product.Post("/insert", backend.InsertProduct)
			})
		})

	}

}
