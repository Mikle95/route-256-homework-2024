package repository

import (
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

func (s *StockStorage) GetStock(sku model.SKU) (model.Stock, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.storage[sku], nil
}

func (s *StockStorage) InsertStock(stock model.Stock) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.storage[stock.Sku] = stock
	return nil
}
