package service

import (
	"context"
	"errors"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/loms_service/loms_client"
	product "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/repository"
)

var _ CartRepository = (*repository.UserCart)(nil)

type CartRepository interface {
	AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error)
	GetItems(ctx context.Context, userId model.UID) ([]model.CartItem, error)
	DeleteItem(ctx context.Context, userId model.UID, sku model.Sku) error
	DeleteCart(ctx context.Context, userId model.UID) error
}

var _ ProductService = (*product.ProductServiceStuct)(nil)

type ProductService interface {
	GetProduct(ctx context.Context, sku model.Sku) (*domain.Item, error)
}

var _ ILOMSService = (*loms_client.Client)(nil)

type ILOMSService interface {
	Checkout(context.Context, domain.Order) (domain.OID, error)
	StocksInfo(context.Context, domain.Sku) (uint64, error)
}

type CartService struct {
	repository     CartRepository
	productService ProductService
	lomsService    ILOMSService
}

func NewCartService(repository CartRepository, ps ProductService, ls ILOMSService) *CartService {
	return &CartService{repository: repository, productService: ps, lomsService: ls}
}

func (s *CartService) AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error) {
	_, err := s.productService.GetProduct(ctx, item.SKU)
	if err != nil {
		return model.CartItem{}, err
	}

	count, err := s.lomsService.StocksInfo(ctx, item.SKU)
	if err != nil {
		return model.CartItem{}, err
	}

	if count < uint64(item.Count) {
		return model.CartItem{}, errors.New("item count > stock total_count")
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

func (s *CartService) Checkout(ctx context.Context, userId model.UID) (domain.OID, error) {
	cart, err := s.repository.GetItems(ctx, userId)
	if err != nil {
		return -1, err
	}

	order_id, err := s.lomsService.Checkout(ctx, repack_cart_order(userId, cart))
	if err != nil {
		return order_id, err
	}

	return order_id, s.repository.DeleteCart(ctx, userId)
}

func repack_cart_order(user_id domain.UID, cart []model.CartItem) domain.Order {
	result := domain.Order{User_id: user_id, Items: make([]domain.OrderItem, 0)}

	for _, item := range cart {
		result.Items = append(result.Items, domain.OrderItem{
			Sku:   item.SKU,
			Count: uint32(item.Count),
		})
	}
	return result
}
