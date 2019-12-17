package fronted

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"seckshop/models"
	"seckshop/services"
	"seckshop/utils"
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
	session := utils.UserSess.Start(ctx)
	userId, err := session.GetInt64("userid")
	if err != nil {
		result.Msg = err.Error()
		result.Code = 500
		result.Data = nil
		ctx.JSON(&result)
		return
	}
	fmt.Println("userid=", userId)
	_, _ = ctx.JSON(p.Service.GetProduct(productId))
}
