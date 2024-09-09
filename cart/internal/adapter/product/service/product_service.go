package product

import (
	"context"
	"errors"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, sku domain.Sku) (*domain.Item, error)
}

type ProductServiceStuct struct {
	productClient ProductRepository
}

func NewProductService(productClient ProductRepository) *ProductServiceStuct {
	return &ProductServiceStuct{productClient: productClient}
}

func (p *ProductServiceStuct) GetProduct(ctx context.Context, sku domain.Sku) (*domain.Item, error) {
	item, err := p.productClient.GetProduct(ctx, sku)
	if err == nil && item.Name == "" {
		return nil, errors.New("sku does not exist")
	}
	return item, err
}
