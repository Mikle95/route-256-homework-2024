package model

type UID = int64
type Sku = int64

type CartItem struct {
	SKU    Sku
	UserId UID
	Count  int16
}
