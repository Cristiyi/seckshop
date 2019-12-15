package models

import "time"

type Product struct {
	ID					int64   `xorm:"id not null pk autoincr int"`
	ProductName  		string `xorm:"varchar(200)"`
	ProductNum   		int8  `xorm:"count int index"`
	ProductFirstImage 	string `xorm:"first_image varchar(255)"`
	ProductImages       string `xorm:"images varchar(255)"`
	ProductUrl   		string `xorm:"product_url varchar(255)"`
	CreatedAt			time.Time `xorm:"created"`
	UpdatedAt 			time.Time `xorm:"updated"`
	DeletedAt 			time.Time `xorm:"deleted"`
}
