package domain

type OID = int64
type OICount = uint32

type Order struct {
	User_id UID
	Items   []OrderItem
}

type OrderItem struct {
	Sku   Sku
	Count OICount
}
