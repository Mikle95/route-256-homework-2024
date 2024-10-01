package loms_service

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

func (s *LOMSService) OrderCreate(ctx context.Context, order model.Order) (model.OID, error) {
	id, err := s.orderS.Create(ctx, order)
	if err != nil {
		return id, err
	}

	err = s.stockS.Reserve(ctx, order.Items)
	if err != nil {
		err1 := s.orderS.SetStatus(ctx, id, model.STATUS_FAIL)
		if err1 != nil {
			panic(err1)
		}
		return id, err
	}

	err = s.orderS.SetStatus(ctx, id, model.STATUS_WAIT)
	if err != nil {
		panic(err)
	}

	return id, nil
}
