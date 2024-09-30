package server

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LOMSServer) OrderPay(ctx context.Context, in *loms.OrderIDRequest) (*loms.EmptyResponse, error) {
	err := s.impl.OrderPay(ctx, in.Info.OrderID)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	return &loms.EmptyResponse{}, nil
}

func (s *LOMSServer) OrderCancel(ctx context.Context, in *loms.OrderIDRequest) (*loms.EmptyResponse, error) {
	err := s.impl.OrderCancel(ctx, in.Info.OrderID)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	return &loms.EmptyResponse{}, nil
}
