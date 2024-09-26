package repository_sqlc

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type OrderStorage struct {
	// q    Querier
	conn *pgx.Conn
}

func NewOrderStorage(conn *pgx.Conn) *OrderStorage {
	return &OrderStorage{
		// q:    New(conn),
		conn: conn,
	}
}

func (s *OrderStorage) GetOrder(_ context.Context, id model.OID) (model.Order, error) {
	panic(errors.New("No impl"))
}

func (s *OrderStorage) AddOrder(_ context.Context, order model.Order) (model.OID, error) {
	panic(errors.New("No impl"))
}

func (s *OrderStorage) ChangeOrder(_ context.Context, id model.OID, order model.Order) error {
	panic(errors.New("No impl"))
}
