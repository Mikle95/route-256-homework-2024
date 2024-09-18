package model

type SKU = uint32
type COUNT = uint64

type Stock struct {
	Sku         SKU
	Total_count COUNT
	Reserved    COUNT
}
