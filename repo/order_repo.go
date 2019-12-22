/**
 * @Author: cristi
 * @Description:
 * @File:  order_repo.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package repo

import (
	"github.com/spf13/cast"
	"seckshop/models"
	_ "seckshop/models"
)

type OrderRepo interface {
	Insert(m map[string]interface{}) (succ bool, order models.Order, msg string)
}

func NewOrderRepo() OrderRepo {
	return &orderRepo{}
}

type orderRepo struct {
}

func (o orderRepo) Insert(m map[string]interface{}) (succ bool, order models.Order, msg string) {
	model := new(models.Order)
	model.UserId = cast.ToInt64(m["user_id"])
	model.ProductId = cast.ToInt64(m["product_id"])
	model.OrderStatus = 1
	model.ProductNum = cast.ToInt32(m["product_num"])
	hasProduct, product, _ := NewProductRepository().GetProduct(model.ProductId)
	if !hasProduct {
		return false, order, "商品不存在"
	}
	model.OrderPrice = product.ProductPrice * cast.ToFloat64(model.ProductNum)
	_, err := engine.Insert(model)
	if err != nil {
		panic("insert error")
	}
	return true, *model, "成功"
}

