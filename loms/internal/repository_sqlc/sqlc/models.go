// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

type Item struct {
	Sku     int32
	OrderID int32
	Count   int32
}

type Stock struct {
	Sku        int32
	TotalCount int32
	Reserved   int32
}

type UserOrder struct {
	OrderID       int32
	UserID        int32
	CurrentStatus string
}
