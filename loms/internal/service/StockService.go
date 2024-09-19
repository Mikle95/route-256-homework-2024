package service

import (
	"errors"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type IStockRepo interface {
	GetStock(model.SKU) (model.Stock, error)
	InsertStock(model.Stock) error
}

type StockService struct {
	stockRepo IStockRepo
	mtx       sync.RWMutex
}

func NewStockService(stockRepo IStockRepo) *StockService {
	return &StockService{stockRepo: stockRepo, mtx: sync.RWMutex{}}
}

func (s *StockService) AddStock(stock model.Stock) error {
	return s.stockRepo.InsertStock(stock)
}

func (s *StockService) GetStockBySKU(sku model.SKU) (model.Stock, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.stockRepo.GetStock(sku)
}

func (s *StockService) Reserve(sku model.SKU, count model.COUNT) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	stock, err := s.stockRepo.GetStock(sku)
	if err != nil {
		return err
	}

	if stock.Reserved+count > stock.Total_count {
		return errors.New("can't reserve more than total_count")
	}

	stock.Reserved += count
	s.stockRepo.InsertStock(stock)

	return nil
}

func (s *StockService) ReserveRemove(sku model.SKU, count model.COUNT) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	stock, err := s.stockRepo.GetStock(sku)
	if err != nil {
		return err
	}

	if stock.Reserved-count < 0 {
		return errors.New("can't remove more than reserved")
	}

	stock.Reserved -= count
	stock.Total_count -= count
	s.stockRepo.InsertStock(stock)

	return nil
}

func (s *StockService) ReserveCancel(sku model.SKU, count model.COUNT) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	stock, err := s.stockRepo.GetStock(sku)
	if err != nil {
		return err
	}

	if stock.Reserved-count < 0 {
		return errors.New("can't cancel more than reserved")
	}

	stock.Reserved -= count
	s.stockRepo.InsertStock(stock)

	return nil
}
