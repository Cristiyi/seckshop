package datasource

import (
	"github.com/go-redis/redis/v7"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12"
	"seckshop/conf"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     conf.Sysconfig.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := Redis.Ping().Err()
	if err != nil {
		panic("redis error")
	}
	iris.RegisterOnInterrupt(func() {
		Redis.Close()
	})
}