package model

type UID = int64
type ItemCount = uint16
type OID = int64

const (
	STATUS_NEW       = "new"
	STATUS_WAIT      = "awaiting payment"
	STATUS_FAIL      = "failed"
	STATUS_PAYED     = "payed"
	STATUS_CANCELLED = "cancelled"
)

type Order struct {
	User_id UID
	Items   []Item
	Status  string
}

type Item struct {
	sku   SKU
	count ItemCount
}
