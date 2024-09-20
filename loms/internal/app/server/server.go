package server

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
)

var _ ILOMSService = (*loms_service.LOMSService)(nil)

type ILOMSService interface {
	OrderCreate(context.Context, model.Order) (model.OID, error)
	OrderInfo(context.Context, model.OID) (model.Order, error)
	OrderPay(context.Context, model.OID) error
	OrderCancel(context.Context, model.OID) error
	StocksInfo(context.Context, model.SKU) (model.COUNT, error)
}

type LOMSServer struct {
	loms.UnimplementedLOMSServer
	impl ILOMSService
}

func NewLOMSServer(impl ILOMSService) *LOMSServer {
	return &LOMSServer{impl: impl}
}
