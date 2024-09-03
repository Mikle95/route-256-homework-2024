package model

type Item struct {
	SKU   Sku
	Name  string
	Price uint32
}

type ItemInfo struct {
	SKU        Sku
	Name       string
	Price      uint32
	Count      uint16
	TotalPrice uint32
}
