/**
 * @Author: cristi
 * @Description:
 * @File:  order_repo.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package repo

import _"seckshop/models"

type OrderRepo interface {

}

func NewOrderRepo() OrderRepo {
	return &orderRepo{}
}

type orderRepo struct {

}