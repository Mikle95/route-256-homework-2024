package loms_service

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

func (s *LOMSService) OrderPay(ctx context.Context, id model.OID) error {
	order, err := s.orderS.GetById(ctx, id)
	if err != nil {
		return err
	}

	if order.Status != model.STATUS_WAIT {
		return fmt.Errorf("wrong order status: %v", order.Status)
	}

	err = s.stockS.ReserveRemove(ctx, order.Items)
	if err != nil {
		panic(err)
	}

	err = s.orderS.SetStatus(ctx, id, model.STATUS_PAYED)
	if err != nil {
		panic(err)
	}

	return nil
}

func (s *LOMSService) OrderCancel(ctx context.Context, id model.OID) error {
	order, err := s.orderS.GetById(ctx, id)
	if err != nil {
		return err
	}

	if order.Status != model.STATUS_WAIT {
		return fmt.Errorf("wrong order status: %v", order.Status)
	}

	err = s.stockS.ReserveCancel(ctx, order.Items)
	if err != nil {
		panic(err)
	}

	err = s.orderS.SetStatus(ctx, id, model.STATUS_CANCELLED)
	if err != nil {
		panic(err)
	}

	return nil
}
