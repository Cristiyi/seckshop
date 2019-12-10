package services

import (
	"seckshop/models"
	"seckshop/repo"
)

type ProductService interface {
	List(m map[string]interface{}) (result models.Result)
	Insert(m map[string]interface{}) (result models.Result)
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

