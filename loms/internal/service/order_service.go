package service

import (
	"context"
	"errors"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository"
)

var _ IOrderRepo = (*repository.OrderStorage)(nil)

type IOrderRepo interface {
	GetOrder(context.Context, model.OID) (model.Order, error)
	AddOrder(context.Context, model.Order) (model.OID, error)
	ChangeOrder(context.Context, model.OID, model.Order) error
}

type OrderService struct {
	orderRepo IOrderRepo
	mtx       sync.RWMutex
}

func NewOrderService(orderRepo IOrderRepo) *OrderService {
	return &OrderService{orderRepo: orderRepo, mtx: sync.RWMutex{}}
}

func (s *OrderService) Create(ctx context.Context, order model.Order) (model.OID, error) {
	order.Status = model.STATUS_NEW
	if len(order.Items) == 0 {
		return -1, errors.New("empty order")
	}
	return s.orderRepo.AddOrder(ctx, order)
}

func (s *OrderService) SetStatus(ctx context.Context, id model.OID, status string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	order, err := s.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return err
	}

	order.Status = status
	s.orderRepo.ChangeOrder(ctx, id, order)
	return nil
}

func (s *OrderService) GetById(ctx context.Context, id model.OID) (model.Order, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.orderRepo.GetOrder(ctx, id)
}
