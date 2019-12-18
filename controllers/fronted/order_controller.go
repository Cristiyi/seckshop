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
		result.Msg = "id不能为空"
		result.Code = 500
		result.Data = nil
		ctx.JSON(&result)
		return
	}
}
