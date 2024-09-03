package server

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type CartService interface {
	AddItem(ctx context.Context, item model.CartItem) error
	GetItems(ctx context.Context, userId model.UID) ([]model.ItemInfo, error)
	DeleteItem(ctx context.Context, userId model.UID, sku model.Sku) error
	DeleteCart(ctx context.Context, userId model.UID) error
}

type CartServer struct {
	cartService CartService
}

func NewCartServer(cartService CartService) *CartServer {
	return &CartServer{cartService: cartService}
}
