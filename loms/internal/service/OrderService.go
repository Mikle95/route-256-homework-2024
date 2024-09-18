package service

import (
	"errors"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type IOrderRepo interface {
	GetOrder(model.OID) (model.Order, error)
	AddOrder(model.Order) (model.OID, error)
	changeOrder(id model.OID, order model.Order) error
}

type IStockService interface {
}

type OrderService struct {
	orderRepo IOrderRepo
}

func NewOrderService(orderRepo IOrderRepo, stockServ IStockService) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) Create(order model.Order) (model.OID, error) {
	order.Status = model.STATUS_NEW
	if len(order.Items) == 0 {
		return -1, errors.New("empty order")
	}
	return s.orderRepo.AddOrder(order)
}
