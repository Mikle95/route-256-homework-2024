package server

import (
	"context"
	"errors"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LOMSServer) OrderInfo(ctx context.Context, in *loms.OrderId) (*loms.OrderInfoResponse, error) {
	order, err := s.impl.OrderInfo(ctx, in.OrderId)
	if errors.Is(err, model.ErrorNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return repack_OrderInfoResponse(order), nil
}

func repack_OrderInfoResponse(order model.Order) *loms.OrderInfoResponse {
	result := loms.OrderInfoResponse{User: order.User_id, Items: make([]*loms.Item, 0)}

	for _, item := range order.Items {
		result.Items = append(result.Items, &loms.Item{
			Sku:   item.Sku,
			Count: uint32(item.Count),
		})
	}

	result.Status = order.Status
	return &result
}
