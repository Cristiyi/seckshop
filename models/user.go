package models

import "time"

type User struct {
	ID				int64 `xorm:"id not null pk autoincr int"`
	Tel				string `xorm:"tel varchar(20)"`
	NickName		string `xorm:"nickname varchar(50)"`
	UserName		string `xorm:"username varchar(50)" index`
	HashPassword    string `xorm:"password varchar(255)"`
	CreatedAt		time.Time `xorm:"created"`
	UpdatedAt 		time.Time `xorm:"updated"`
	DeletedAt 		time.Time `xorm:"deleted"`
}
