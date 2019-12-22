package api

import (
	"github.com/kataras/iris/v12"
	"seckshop/services"
)

type ApiController struct {
	productService services.ProductService
}

func NewApiController() *ApiController {
	return &ApiController{services.NewProductService()}
}

//缓存所有秒杀商品库存
func SetAllSeckCount(ctx iris.Context) {
	c := NewApiController()
	_, _ = ctx.JSON(c.productService.SetAllSeckCount())
}
