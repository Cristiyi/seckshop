package backend

import (
	"github.com/kataras/iris/v12"
	"seckshop/conf"
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

func ProductList(ctx iris.Context) {
	r := ctx.Request()
	result := new(models.Result)
	m := make(map[string]interface{})
	page := r.PostFormValue("page")
	size := r.PostFormValue("size")
	if page == "" {
		result.Msg = "page不能为空"
		result.Code = -1
		result.Data = nil
		ctx.JSON(&result)
		return
	}
	if size == "" {
		result.Code = -1
		result.Msg = "size不能为空"
		result.Data = nil
		ctx.JSON(&result)
		return
	}
	m["page"] = page
	m["size"] = size
	p := NewProductController()
	_, _ = ctx.JSON(p.Service.List(m))
}

func InsertProduct(ctx iris.Context) {

	r := ctx.Request()
	result := new(models.Result)
	productName := r.PostFormValue("product_name")
	if productName == "" {
		result.Msg = "请传入商品名"
		result.Code = 500
		ctx.JSON(&result)
		return
	}
	productNum := r.PostFormValue("product_num")
	if productNum == "" {
		result.Msg = "请传入商品数量"
		result.Code = 500
		ctx.JSON(&result)
		return
	}

	//处理商品图片（多图）
	imageParams := r.MultipartForm
	images := imageParams.File["images"]
	if images == nil {
		result.Msg = "请传入商品图片"
		result.Code = 500
		ctx.JSON(&result)
		return
	}
	uploadDest := conf.ProductUploadDest
	maxSize := conf.MaxImageUploadSize
	imageEerr := ctx.Request().ParseMultipartForm(maxSize)
	if imageEerr != nil {
		result.Msg = "图片文件过大"
		result.Code = 500
		ctx.JSON(&result)
		return
	}
	m := make(map[string]interface{})
	m["images"] = utils.UploadFiles(images, uploadDest)
	if m["images"] == "" {
		result.Msg = "商品图片上传错误"
		result.Code = 500
		ctx.JSON(&result)
		return
	}

	//处理商品主图（单图）
	file, info, firstImageErr := ctx.FormFile("first_image")
	if firstImageErr != nil {
		result.Msg = "请上传主图"
		result.Code = 500
		ctx.JSON(&result)
		return
	}
	if info.Size > maxSize {
		result.Msg = "图片文件过大"
		result.Code = 500
		ctx.JSON(&result)
		return
	}
	m["first_image"] = utils.UploadFile(file, info, uploadDest)
	if m["first_image"] == "" {
		result.Msg = "商品主图上传错误"
		result.Code = 500
		ctx.JSON(&result)
		return
	}

	if r.PostFormValue("product_url") != "" {
		m["product_url"] = r.PostFormValue("product_url")
	}
	m["product_name"] = productName
	m["product_num"] = productNum
	p := NewProductController()
	_, _ = ctx.JSON(p.Service.Insert(m))

}
