package datasource

import (
	"seckshop/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"strings"
	"seckshop/models"
)

var engine *xorm.Engine

func GetDB() *xorm.Engine {
	return engine
}

func init() {
	path := strings.Join([]string{conf.Sysconfig.DBUserName, ":", conf.Sysconfig.DBPassword, "@(", conf.Sysconfig.DBIp, ":", conf.Sysconfig.DBPort, ")/", conf.Sysconfig.DBName, "?charset=utf8&parseTime=true"}, "")
	var err error
	engine, err = xorm.NewEngine("mysql", path)
	if err != nil {
		panic("database error")
	}
	engine.ShowSQL(true)

	//创建表

	err = engine.Sync2(new(models.Product), new(models.Order), new(models.User))

	if err != nil {
		panic(err)
	}

}

