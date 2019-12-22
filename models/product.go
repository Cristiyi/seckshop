package models

import "time"

type Product struct {
	ID					int64   `xorm:"id not null pk autoincr int"`
	ProductName  		string `xorm:"varchar(200)"`
	ProductNum   		int32   `xorm:"count int index"`
	ProductPrice		float64 `xorm:"price decimal(6,2)"`
	ProductFirstImage 	string `xorm:"first_image varchar(255)"`
	ProductImages       string `xorm:"images varchar(255)"`
	ProductUrl   		string `xorm:"product_url varchar(255)"`
	CreatedAt			time.Time `xorm:"created"`
	UpdatedAt 			time.Time `xorm:"updated"`
	DeletedAt 			time.Time `xorm:"deleted"`
}
