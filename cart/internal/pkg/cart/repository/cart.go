package repository

import (
	"context"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type Storage = map[model.Sku]model.CartItem

type Cart struct {
	storage Storage
	mtx     sync.RWMutex
}

func NewCart() *Cart {
	return &Cart{storage: make(Storage), mtx: sync.RWMutex{}}
}

func (c *Cart) AddItem(_ context.Context, item model.CartItem) model.CartItem {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cartItem, exists := c.storage[item.SKU]
	if exists {
		item.Count += cartItem.Count
	}
	c.storage[item.SKU] = item
	return item
}

func (c *Cart) GetItems(_ context.Context) []model.CartItem {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	out := make([]model.CartItem, 0, len(c.storage))
	for _, val := range c.storage {
		out = append(out, val)
	}
	return out
}

func (c *Cart) DeleteItem(_ context.Context, sku model.Sku) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, exists := c.storage[sku]
	if exists {
		delete(c.storage, sku)
	}
}
