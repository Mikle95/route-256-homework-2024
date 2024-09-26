package repository_sqlc

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type StockStorage struct {
	// q    Querier
	conn *pgx.Conn
}

func NewStockStorage(conn *pgx.Conn) *StockStorage {
	return &StockStorage{
		// q:    New(conn),
		conn: conn,
	}
}

func (s *StockStorage) GetStock(_ context.Context, sku model.SKU) (model.Stock, error) {
	panic(errors.New("No impl"))
}

func (s *StockStorage) InsertStock(_ context.Context, stock model.Stock) error {
	panic(errors.New("No impl"))
}
