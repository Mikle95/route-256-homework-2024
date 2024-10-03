package repository

import (
	"context"
	"fmt"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type OrderStorage struct {
	storage []model.Order
	mtx     sync.RWMutex
}

func NewOrderStorage() *OrderStorage {
	return &OrderStorage{storage: make([]model.Order, 0), mtx: sync.RWMutex{}}
}

func (s *OrderStorage) GetOrder(_ context.Context, id model.OID) (model.Order, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if id >= model.OID(len(s.storage)) || id < 0 {
		return model.Order{}, fmt.Errorf("%w: get order: wrong id", model.ErrorNotFound)
	}

	return s.storage[id], nil
}

func (s *OrderStorage) AddOrder(_ context.Context, order model.Order) (model.OID, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.storage = append(s.storage, order)
	return model.OID(len(s.storage) - 1), nil
}

func (s *OrderStorage) ChangeOrder(_ context.Context, id model.OID, order model.Order) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if id >= model.OID(len(s.storage)) || id < 0 {
		return fmt.Errorf("%w: change order: wrong id", model.ErrorNotFound)
	}

	s.storage[id] = order

	return nil
}
