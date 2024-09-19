package repository

import (
	"context"
	"errors"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type StockStorage struct {
	storage map[model.SKU]model.Stock
	mtx     sync.RWMutex
}

func NewStockStorage() *StockStorage {
	return &StockStorage{storage: make(map[model.SKU]model.Stock), mtx: sync.RWMutex{}}
}

func (s *StockStorage) GetStock(_ context.Context, sku model.SKU) (model.Stock, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	stock, flag := s.storage[sku]
	var err error
	if !flag {
		err = errors.New("wrong sku")
	}

	return stock, err
}

func (s *StockStorage) InsertStock(_ context.Context, stock model.Stock) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.storage[stock.Sku] = stock
	return nil
}
