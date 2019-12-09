package backend

import (
	"github.com/kataras/iris"
	"seckshop/models"
	"seckshop/services"
)

type ProductController struct {
	Ctx iris.Context
	Service services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{Service:services.NewProductService()}
}

func (g *ProductController) PostList()(result models.Result)  {
	r := g.Ctx.Request()
	m:=make(map[string]interface{})
	page := r.PostFormValue("page")
	size := r.PostFormValue("size")
	if page == "" {
		result.Msg = "page不能为空"
		result.Code = -1
		return
	}
	if size == "" {
		result.Code = -1
		result.Msg = "size不能为空"
		return
	}
	m["page"] = page
	m["size"] = size
	return g.Service.List(m)
}
