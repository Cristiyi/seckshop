package fronted

import (
	"github.com/kataras/iris/v12"
	"seckshop/services"
	"seckshop/models"
	"github.com/spf13/cast"
)

type ProductController struct {
	Service services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{Service:services.NewProductService()}
}

func Detail(ctx iris.Context) () {
	result := new(models.Result)
	p := NewProductController()
	r := ctx.Request()
	if r.PostFormValue("id") == "" {
		result.Msg = "id不能为空"
		result.Code = 500
		result.Data = nil
		ctx.JSON(&result)
		return
	}
	productId := cast.ToInt64(r.PostFormValue("id"))
	//return p.Service.GetProduct(productId)
	_, _ = ctx.JSON(p.Service.GetProduct(productId))
}
