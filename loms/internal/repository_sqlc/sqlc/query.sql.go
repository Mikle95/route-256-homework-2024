// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
)

const addUpdateStock = `-- name: AddUpdateStock :exec
INSERT INTO stock (sku, total_count, reserved)
VALUES ($1, $2, $3)
ON CONFLICT(sku) 
DO UPDATE SET total_count = $2, reserved = $3
`

type AddUpdateStockParams struct {
	Sku        int32
	TotalCount int32
	Reserved   int32
}

func (q *Queries) AddUpdateStock(ctx context.Context, arg *AddUpdateStockParams) error {
	_, err := q.db.Exec(ctx, addUpdateStock, arg.Sku, arg.TotalCount, arg.Reserved)
	return err
}

const getStock = `-- name: GetStock :one
SELECT sku, total_count, reserved FROM stock
WHERE sku = $1
`

func (q *Queries) GetStock(ctx context.Context, sku int32) (*Stock, error) {
	row := q.db.QueryRow(ctx, getStock, sku)
	var i Stock
	err := row.Scan(&i.Sku, &i.TotalCount, &i.Reserved)
	return &i, err
}

const insertItem = `-- name: InsertItem :exec
INSERT INTO item (sku, order_id, count)
VALUES ($1, $2, $3)
`

type InsertItemParams struct {
	Sku     int32
	OrderID int32
	Count   int32
}

func (q *Queries) InsertItem(ctx context.Context, arg *InsertItemParams) error {
	_, err := q.db.Exec(ctx, insertItem, arg.Sku, arg.OrderID, arg.Count)
	return err
}

const insertOrder = `-- name: InsertOrder :one
INSERT INTO user_order (user_id, current_status)
VALUES ($1, $2)
RETURNING order_id
`

type InsertOrderParams struct {
	UserID        int32
	CurrentStatus string
}

func (q *Queries) InsertOrder(ctx context.Context, arg *InsertOrderParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertOrder, arg.UserID, arg.CurrentStatus)
	var order_id int32
	err := row.Scan(&order_id)
	return order_id, err
}

const selectItem = `-- name: SelectItem :many
SELECT sku, order_id, count FROM item
WHERE order_id = $1
`

func (q *Queries) SelectItem(ctx context.Context, orderID int32) ([]*Item, error) {
	rows, err := q.db.Query(ctx, selectItem, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.Sku, &i.OrderID, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectOrder = `-- name: SelectOrder :one
SELECT order_id, user_id, current_status FROM user_order
WHERE order_id = $1
`

func (q *Queries) SelectOrder(ctx context.Context, orderID int32) (*UserOrder, error) {
	row := q.db.QueryRow(ctx, selectOrder, orderID)
	var i UserOrder
	err := row.Scan(&i.OrderID, &i.UserID, &i.CurrentStatus)
	return &i, err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :exec
UPDATE user_order
SET current_status = $1
WHERE order_id = $2
`

type UpdateOrderStatusParams struct {
	CurrentStatus string
	OrderID       int32
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg *UpdateOrderStatusParams) error {
	_, err := q.db.Exec(ctx, updateOrderStatus, arg.CurrentStatus, arg.OrderID)
	return err
}
