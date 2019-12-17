package services

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"seckshop/datasource"
	"seckshop/models"
	"seckshop/repo"
)

type ProductService interface {
	List(m map[string]interface{}) (result models.Result)
	Insert(m map[string]interface{}) (result models.Result)
	GetProduct(productId int64) (result models.Result)
}

type productService struct {
	Repository repo.ProductRepository
}

func NewProductService() ProductService {
	return &productService{Repository:repo.NewProductRepository()}
}

func (p *productService) List(m map[string]interface{}) (result models.Result){
	total, productList := p.Repository.List(m)
	maps := make(map[string]interface{}, 2)
	maps["total"] = total
	maps["productList"] = productList
	result.Data = maps
	result.Code = 200
	result.Msg = "success"
	return
}

func (p *productService) Insert(m map[string]interface{}) (result models.Result) {
	goodId, err := p.Repository.Insert(m)
	if err != nil {
		panic("insert error")
	}
	result.Data = goodId
	result.Code = 200
	result.Msg = "添加成功"
	return
}

func (p *productService) GetProduct(productId int64) (result models.Result) {
	hasProduct, product, err := p.Repository.GetProduct(productId)
	if err != nil {
		panic("error")
	}
	if hasProduct == false {
		result.Data = nil
		result.Code = 205
		result.Msg = "暂无数据"
		return
	}
	result.Data = product
	result.Code = 200
	result.Msg = "执行成功"
	productJson, _ := json.Marshal(product)
	rErr := datasource.Redis.HSet("SECKPRODUCT", cast.ToString(product.ID), productJson).Err()
	if rErr != nil {
		fmt.Println(rErr)
	}
	return
}

