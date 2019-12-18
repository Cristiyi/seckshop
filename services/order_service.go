package services

import "seckshop/repo"

type OrderService interface {

}

type orderService struct {
	Repository repo.OrderRepo
}

func NewOrderService() OrderService {
	return &orderService{Repository:repo.NewOrderRepo()}
}