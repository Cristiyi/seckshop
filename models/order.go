/**
 * @Author: cristi
 * @Description:
 * @File:  order.go
 * @Version: 1.0.0
 * @Date: 2019/12/10 0010 14:08
 */

package models

import "time"

//order实体
type Order struct {
	ID				int64 `xorm:"id not null pk autoincr int"`
	UserId			int64 `xorm:"user_id int index"`
	ProductId		int64 `xorm:"product_id int index"`
	ProductNum      int8  `xorm:"product_num int"`
	OrderPrice      float64 `xorm:"price" decimal(6,2)`
	OrderStatus		int8  `xorm:"status int index"`
	CreatedAt		time.Time `xorm:"created"`
	UpdatedAt 		time.Time `xorm:"updated"`
	DeletedAt 		time.Time `xorm:"deleted"`
}
