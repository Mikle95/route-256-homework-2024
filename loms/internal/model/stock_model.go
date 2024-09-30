package model

type SKU = uint32
type COUNT = uint64

type Stock struct {
	Sku         SKU   `json:"sku"`
	Total_count COUNT `json:"total_count"`
	Reserved    COUNT `json:"reserved"`
}
