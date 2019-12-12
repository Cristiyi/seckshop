package backend

import (
	"github.com/kataras/iris/v12"
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
	m := make(map[string]interface{})
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

func (g *ProductController) PostInsert(result models.Result) {
	r := g.Ctx.Request()
	productName := r.PostFormValue("product_name")
	if productName == "" {
		result.Msg = "请传入商品名"
		result.Code = 500
		return
	}
	productNum := r.PostFormValue("product_num")
	if productNum == "" {
		result.Msg = "请传入商品数量"
		result.Code = 500
		return
	}
	firstImage := r.PostFormValue("first_image")
	if firstImage == "" {
		result.Msg = "请传入商品图片"
		result.Code = 500
		return
	}
	m := make(map[string]interface{})
	if r.PostFormValue("product_url") != "" {
		m["product_url"] = r.PostFormValue("product_url")
	}
	m["product_name"] = productName
	m["product_num"] = productNum
	m["first_image"] = firstImage
}
