package repository

import (
	"errors"
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

func (s *OrderStorage) GetOrder(id model.OID) (model.Order, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if id >= model.OID(len(s.storage)) {
		return model.Order{}, errors.New("wrong id")
	}

	return s.storage[id], nil
}

func (s *OrderStorage) AddOrder(order model.Order) (model.OID, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.storage = append(s.storage, order)
	return model.OID(len(s.storage) - 1), nil
}

func (s *OrderStorage) changeOrder(id model.OID, order model.Order) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if id <= model.OID(len(s.storage)) {
		return errors.New("change order: wrong id!")
	}

	s.storage[id] = order

	return nil
}
