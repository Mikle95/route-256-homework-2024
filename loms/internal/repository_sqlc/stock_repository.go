package repository_sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository_sqlc/sqlc"
)

type StockStorage struct {
	q    sqlc.Querier
	conn *pgx.Conn
}

func NewStockStorage(conn *pgx.Conn) *StockStorage {
	return &StockStorage{
		q:    sqlc.New(conn),
		conn: conn,
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
