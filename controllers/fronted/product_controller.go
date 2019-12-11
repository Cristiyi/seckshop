package fronted

import (
	"github.com/kataras/iris"
	"seckshop/services"
	"seckshop/models"
	"github.com/spf13/cast"
)

type ProductController struct {
	Ctx		iris.Context
	Service services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{Service:services.NewProductService()}
}

func (g *ProductController) PostGet() (result models.Result) {
	r := g.Ctx.Request()
	if r.PostFormValue("id") == "" {
		result.Msg = "id不能为空"
		result.Code = -1
		return
	}
	productId := cast.ToInt64(r.PostFormValue("id"))
	return g.Service.GetProduct(productId)
}
