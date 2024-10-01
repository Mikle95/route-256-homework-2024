package server

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LOMSServer) OrderCreate(ctx context.Context, in *loms.OrderCreateRequest) (*loms.OrderID, error) {
	id, err := s.impl.OrderCreate(ctx, repack(in))
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	return &loms.OrderID{OrderID: id}, nil
}

func repack(order *loms.OrderCreateRequest) model.Order {
	result := model.Order{User_id: order.Info.User, Items: make([]model.Item, 0)}

	for _, item := range order.Info.Items {
		result.Items = append(result.Items, model.Item{
			Sku:   item.Sku,
			Count: model.ItemCount(item.Count),
		})
	}
	return result
}
