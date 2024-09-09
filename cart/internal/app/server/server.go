package server

import (
	"context"

	gody "github.com/guiferpa/gody/v2"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type CartService interface {
	AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error)
	GetItems(ctx context.Context, userId model.UID) (*model.UserCartInfo, error)
	DeleteItem(ctx context.Context, userId model.UID, sku model.Sku) error
	DeleteCart(ctx context.Context, userId model.UID) error
}

type Response struct {
	Message string `json:"message"`
}

type CartServer struct {
	cartService CartService
	validator   *gody.Validator
}

func NewCartServer(cartService CartService, v *gody.Validator) *CartServer {
	return &CartServer{cartService: cartService, validator: v}
}
