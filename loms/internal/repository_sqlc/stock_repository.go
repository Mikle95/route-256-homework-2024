package repository_sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository_sqlc/sqlc"
)

type StockStorage struct {
	q    sqlc.Querier
	pool *pgxpool.Pool
}

func NewStockStorage(pool *pgxpool.Pool) *StockStorage {
	return &StockStorage{
		q:    sqlc.New(pool),
		pool: pool,
	}
}

func (s *StockStorage) GetStock(ctx context.Context, sku model.SKU) (model.Stock, error) {
	stock, err := s.q.GetStock(ctx, int32(sku))
	if err != nil {
		return model.Stock{}, err
	}
	return model.Stock{Sku: sku, Total_count: uint64(stock.TotalCount), Reserved: uint64(stock.Reserved)}, nil
}

func (s *StockStorage) InsertStock(ctx context.Context, stock model.Stock) error {
	return s.q.AddUpdateStock(ctx, &sqlc.AddUpdateStockParams{Sku: int32(stock.Sku), TotalCount: int32(stock.Total_count), Reserved: int32(stock.Reserved)})
}
