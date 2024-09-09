package domain

type UID = int64
type Sku = int64
type Count = uint16
type Price = uint32

type Item struct {
	SKU   Sku
	Name  string
	Price Price
}
