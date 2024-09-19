package service

import (
	"context"
	"errors"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository"
)

var _ IStockRepo = (*repository.StockStorage)(nil)

type IStockRepo interface {
	GetStock(context.Context, model.SKU) (model.Stock, error)
	InsertStock(context.Context, model.Stock) error
}

type StockService struct {
	stockRepo IStockRepo
	mtx       sync.RWMutex
}

func NewStockService(stockRepo IStockRepo) *StockService {
	return &StockService{stockRepo: stockRepo, mtx: sync.RWMutex{}}
}

func (s *StockService) AddStock(ctx context.Context, stock model.Stock) error {
	return s.stockRepo.InsertStock(ctx, stock)
}

func (s *StockService) GetStockBySKU(ctx context.Context, sku model.SKU) (model.Stock, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.stockRepo.GetStock(ctx, sku)
}

func (s *StockService) Reserve(ctx context.Context, items []model.Item) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, item := range items {
		stock, err := s.stockRepo.GetStock(ctx, item.Sku)
		if err != nil {
			return err
		}

		if stock.Reserved+uint64(item.Count) > stock.Total_count {
			return errors.New("can't reserve more than total_count")
		}
	}

	for _, item := range items {
		stock, err := s.stockRepo.GetStock(ctx, item.Sku)
		if err != nil {
			panic(err)
		}
		stock.Reserved += uint64(item.Count)
		err = s.stockRepo.InsertStock(ctx, stock)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (s *StockService) ReserveRemove(ctx context.Context, items []model.Item) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, item := range items {
		stock, err := s.stockRepo.GetStock(ctx, item.Sku)
		if err != nil {
			return err
		}

		if stock.Reserved-uint64(item.Count) < 0 {
			return errors.New("can't remove more than reserved")
		}
	}

	for _, item := range items {
		stock, err := s.stockRepo.GetStock(ctx, item.Sku)
		if err != nil {
			panic(err)
		}
		stock.Reserved -= uint64(item.Count)
		stock.Total_count -= uint64(item.Count)
		err = s.stockRepo.InsertStock(ctx, stock)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (s *StockService) ReserveCancel(ctx context.Context, items []model.Item) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, item := range items {
		stock, err := s.stockRepo.GetStock(ctx, item.Sku)
		if err != nil {
			return err
		}

		if stock.Reserved-uint64(item.Count) < 0 {
			return errors.New("can't cancel more than reserved")
		}
	}

	for _, item := range items {
		stock, err := s.stockRepo.GetStock(ctx, item.Sku)
		if err != nil {
			panic(err)
		}
		stock.Reserved -= uint64(item.Count)
		stock.Total_count -= uint64(item.Count)
		err = s.stockRepo.InsertStock(ctx, stock)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
