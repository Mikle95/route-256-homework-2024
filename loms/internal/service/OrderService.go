package service

import (
	"errors"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type IOrderRepo interface {
	GetOrder(model.OID) (model.Order, error)
	AddOrder(model.Order) (model.OID, error)
	ChangeOrder(id model.OID, order model.Order) error
}

type OrderService struct {
	orderRepo IOrderRepo
	mtx       sync.RWMutex
}

func NewOrderService(orderRepo IOrderRepo) *OrderService {
	return &OrderService{orderRepo: orderRepo, mtx: sync.RWMutex{}}
}

func (s *OrderService) Create(order model.Order) (model.OID, error) {
	order.Status = model.STATUS_NEW
	if len(order.Items) == 0 {
		return -1, errors.New("empty order")
	}
	return s.orderRepo.AddOrder(order)
}

func (s *OrderService) SetStatus(id model.OID, status string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	order, err := s.orderRepo.GetOrder(id)
	if err != nil {
		return err
	}

	order.Status = status
	s.orderRepo.ChangeOrder(id, order)
	return nil
}

func (s *OrderService) GetById(id model.OID) (model.Order, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.orderRepo.GetOrder(id)
}
