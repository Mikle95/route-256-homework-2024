package service

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, sku model.Sku) (*model.Item, error)
}

type ProductServiceStuct struct {
	productClient ProductRepository
}

func NewProductService(productClient ProductRepository) *ProductServiceStuct {
	return &ProductServiceStuct{productClient: productClient}
}

// TODO Добавить ретраи 420/429
func (p *ProductServiceStuct) GetProduct(ctx context.Context, sku model.Sku) (*model.Item, error) {
	return p.productClient.GetProduct(ctx, sku)
}
