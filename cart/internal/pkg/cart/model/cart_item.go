package model

type UID = int64
type Sku = int64
type Count = uint16

type CartItem struct {
	SKU    Sku
	UserId UID
	Count  Count
}
