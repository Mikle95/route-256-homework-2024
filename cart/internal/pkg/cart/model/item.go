package model

type Price = uint32

type ItemInfo struct {
	SKU   Sku    `json:"sku_id"`
	Name  string `json:"name"`
	Price Price  `json:"price"`
	Count Count  `json:"count"`
}

type UserCartInfo struct {
	Items      []ItemInfo `json:"items"`
	TotalPrice Price      `json:"total_price"`
}
