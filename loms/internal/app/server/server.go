package server

import (
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
)

type ILOMSService interface {
	OrderCreate(order model.Order) (model.OID, error)
	OrderInfo(id model.OID) (model.Order, error)
	OrderPay(id model.OID) error
	OrderCancel(id model.OID) error
	StocksInfo(sku model.SKU) (model.COUNT, error)
}

type LOMSServer struct {
	loms.UnimplementedLOMSServer
	impl ILOMSService
}

func NewLOMSServer(impl ILOMSService) *LOMSServer {
	return &LOMSServer{impl: impl}
}
