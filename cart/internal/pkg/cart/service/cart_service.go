package service

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type CartRepository interface {
	AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error)
	GetItems(ctx context.Context, userId model.UID) ([]model.CartItem, error)
	DeleteItem(ctx context.Context, userId model.UID, sku model.Sku) error
	DeleteCart(ctx context.Context, userId model.UID) error
}

type ProductService interface {
	GetProduct(ctx context.Context, sku model.Sku) (*domain.Item, error)
}

type CartService struct {
	repository     CartRepository
	productService ProductService
}

func NewCartService(repository CartRepository, ps ProductService) *CartService {
	return &CartService{repository: repository, productService: ps}
}

func (s *CartService) AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error) {
	_, err := s.productService.GetProduct(ctx, item.SKU)
	if err != nil {
		return model.CartItem{}, err
	}
	return s.repository.AddItem(ctx, item)
}

func (s *CartService) GetItems(ctx context.Context, userId model.UID) (*model.UserCartInfo, error) {
	mas, err := s.repository.GetItems(ctx, userId)
	if err != nil {
		return nil, err
	}

	result := model.UserCartInfo{
		Items:      make([]model.ItemInfo, 0, len(mas)),
		TotalPrice: 0,
	}

	for _, val := range mas {
		item, err := s.productService.GetProduct(ctx, val.SKU)
		if err != nil {
			return nil, err
		}
		result.Items = append(result.Items, model.ItemInfo{
			SKU:   val.SKU,
			Name:  item.Name,
			Price: item.Price,
			Count: val.Count,
		})

		result.TotalPrice += item.Price * uint32(val.Count)
	}
	return &result, nil
}

func (s *CartService) DeleteItem(ctx context.Context, userId model.UID, sku model.Sku) error {
	return s.repository.DeleteItem(ctx, userId, sku)
}

func (s *CartService) DeleteCart(ctx context.Context, userId model.UID) error {
	return s.repository.DeleteCart(ctx, userId)
}
