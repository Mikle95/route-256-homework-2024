package repository_sqlc

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository_sqlc/sqlc"
)

type OrderStorage struct {
	q    sqlc.Querier
	conn *pgx.Conn
}

func NewOrderStorage(conn *pgx.Conn) *OrderStorage {
	return &OrderStorage{
		q:    sqlc.New(conn),
		conn: conn,
	}
}

func (s *OrderStorage) GetOrder(ctx context.Context, id model.OID) (model.Order, error) {
	order, err := s.q.SelectOrder(ctx, int32(id))
	if err != nil {
		return model.Order{}, fmt.Errorf("SelectOrder: %w", err)
	}

	items, err := s.q.SelectItem(ctx, int32(id))
	if err != nil {
		return model.Order{}, fmt.Errorf("SelectItem: %w", err)
	}

	result := model.Order{User_id: int64(order.UserID), Status: order.CurrentStatus, Items: make([]model.Item, 0, len(items))}
	for _, item := range items {
		result.Items = append(result.Items, model.Item{Sku: uint32(item.Sku), Count: uint16(item.Count)})
	}

	return result, nil
}

func (s *OrderStorage) AddOrder(ctx context.Context, order model.Order) (id model.OID, err error) {
	err = pgx.BeginFunc(ctx, s.conn, func(tx pgx.Tx) error {
		repo := sqlc.New(tx)
		order_id, err := repo.InsertOrder(ctx, &sqlc.InsertOrderParams{UserID: int32(order.User_id), CurrentStatus: order.Status})
		if err != nil {
			return err
		}

		for _, item := range order.Items {
			err = repo.InsertItem(ctx, &sqlc.InsertItemParams{Sku: int32(item.Sku), OrderID: order_id, Count: int32(item.Count)})
			if err != nil {
				return err
			}
		}
		id = int64(order_id)
		return nil
	})

	return
}

func (s *OrderStorage) ChangeOrder(_ context.Context, id model.OID, order model.Order) error {
	return s.q.UpdateOrderStatus(context.Background(), &sqlc.UpdateOrderStatusParams{OrderID: int32(id), CurrentStatus: order.Status})
}
