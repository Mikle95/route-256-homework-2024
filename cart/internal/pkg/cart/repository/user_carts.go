package repository

import (
	"context"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type UserStorage = map[model.UID]Cart
type UserCart struct {
	storage UserStorage
	mtx     sync.RWMutex
}

func NewUserStorage() *UserCart {
	return &UserCart{storage: make(UserStorage), mtx: sync.RWMutex{}}
}

func (c *UserCart) AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cart, exist := c.storage[item.UserId]
	if !exist {
		cart = *NewCart()
	}
	item = cart.AddItem(ctx, item)
	c.storage[item.UserId] = cart
	return item, nil
}

func (c *UserCart) DeleteItem(ctx context.Context, userId model.UID, sku model.Sku) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cart, exist := c.storage[userId]
	if exist {
		cart.DeleteItem(ctx, sku)
	}
	return nil
}

func (c *UserCart) GetItems(ctx context.Context, userId model.UID) ([]model.CartItem, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	cart, exist := c.storage[userId]
	if exist {
		return cart.GetItems(ctx), nil
	}
	return nil, nil
}

func (c *UserCart) DeleteCart(ctx context.Context, userId model.UID) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, exist := c.storage[userId]
	if exist {
		delete(c.storage, userId)
	}
	return nil
}
