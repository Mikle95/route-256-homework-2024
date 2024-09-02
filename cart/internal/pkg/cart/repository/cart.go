package repository

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type Storage = map[model.Sku]model.CartItem

type Cart struct {
	storage Storage
}

func NewCart(capacity int) *Cart {
	return &Cart{storage: make(Storage)}
}

func (c *Cart) AddItem(_ context.Context, item model.CartItem) model.CartItem {
	cartItem, exists := c.storage[item.SKU]
	if exists {
		item.Count += cartItem.Count
	}
	c.storage[item.SKU] = item
	return item
}

func (c *Cart) GetItems(_ context.Context, sku model.Sku) []model.CartItem {
	out := make([]model.CartItem, 0, len(c.storage))
	for _, val := range c.storage {
		out = append(out, val)
	}
	return out
}

func (c *Cart) DeleteItem(_ context.Context, item model.CartItem) {
	_, exists := c.storage[item.SKU]
	if exists {
		delete(c.storage, item.SKU)
	}
}
