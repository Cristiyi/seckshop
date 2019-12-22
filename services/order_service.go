package services

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
	"seckshop/datasource"
	"seckshop/models"
	"seckshop/repo"
	"time"
)

type OrderService interface {
	Insert(m map[string]interface{})(models.Result)
	CheckRedisCount(productId string) (succ bool, msg string)
}

type orderService struct {
	Repository repo.OrderRepo
}

func NewOrderService() OrderService {
	return &orderService{Repository:repo.NewOrderRepo()}
}

var fn = func(tx *redis.Tx) error {
	// 先查询下当前watch监听的key的值
	v, err := tx.Get("key").Result()
	if err != nil && err != redis.Nil {
		return err
	}
	value := cast.ToInt64(v)
	// 如果key的值没有改变的话，Pipelined函数才会调用成功
	_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
		// 在这里给key设置最新值
		pipe.Set("key", value, time.Hour)
		return nil
	})
	return err
}

func (p *orderService) Insert(m map[string]interface{}) (result models.Result) {
	succ, order, msg := p.Repository.Insert(m)
	if succ {
		result.Data = order
		result.Code = 200
		result.Msg = msg
		return
	} else {
		result.Data = nil
		result.Code = 500
		result.Msg = msg
		return
	}
}

func (p *orderService) CheckRedisCount(productId string) (succ bool, msg string) {

	redisObj := datasource.Redis
	msg = "商品已售完"
	redisResult := redisObj.Get(productId)
	productNum, error := redisResult.Int64()
	if error != nil {
		return false, msg
	}
	if productNum <= 0 {
		return false, msg
	}

	err := redisObj.Watch(fn, productId)

	if err != nil {
		return false, msg
	} else {
		return true, "success"
	}

}