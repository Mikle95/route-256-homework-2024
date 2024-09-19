package loms_service

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service"
)

var _ IStockService = (*service.StockService)(nil)
var _ IOrderService = (*service.OrderService)(nil)

type IStockService interface {
	GetStockBySKU(context.Context, model.SKU) (model.Stock, error)
	Reserve(context.Context, []model.Item) error
	ReserveRemove(context.Context, []model.Item) error
	ReserveCancel(context.Context, []model.Item) error
}

type IOrderService interface {
	Create(context.Context, model.Order) (model.OID, error)
	SetStatus(context.Context, model.OID, string) error
	GetById(context.Context, model.OID) (model.Order, error)
}

type LOMSService struct {
	orderS IOrderService
	stockS IStockService
}

func NewLOMSService(orderS IOrderService, stockS IStockService) *LOMSService {
	return &LOMSService{orderS: orderS, stockS: stockS}
}

func (s *LOMSService) OrderInfo(ctx context.Context, id model.OID) (model.Order, error) {
	return s.orderS.GetById(ctx, id)
}

func (s *LOMSService) StocksInfo(ctx context.Context, sku model.SKU) (model.COUNT, error) {
	stock, err := s.stockS.GetStockBySKU(ctx, sku)
	if err != nil {
		return 0, err
	}
	return stock.Total_count - stock.Reserved, nil
}
