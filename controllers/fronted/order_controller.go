package fronted

import (
	"github.com/kataras/iris/v12"
	"seckshop/models"
	"seckshop/services"
)

type OrderController struct {
	Service services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{Service:services.NewOrderService()}
}

func Seck(ctx iris.Context) {
	r := ctx.Request()
	result := new(models.Result)
	if r.PostFormValue("product_id") == "" {
		result.Msg = "商品id不能为空"
		result.Code = 500
		result.Data = nil
		ctx.JSON(&result)
		return
	}
	productId := r.PostFormValue("product_id")

	c := NewOrderController()
	succ, msg := c.Service.CheckRedisCount(productId)
	if succ {
		result.Msg = "抢购成功，请稍后"
		result.Code = 200
		result.Data = nil
	} else {
		result.Msg = msg
		result.Code = 500
		result.Data = nil
	}
	ctx.JSON(&result)

}
