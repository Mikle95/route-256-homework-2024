package server

import (
	"context"
	"errors"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LOMSServer) OrderPay(ctx context.Context, in *loms.OrderId) (*loms.EmptyResponse, error) {
	err := s.impl.OrderPay(ctx, in.OrderId)
	if errors.Is(err, model.ErrorNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &loms.EmptyResponse{}, nil
}

func (s *LOMSServer) OrderCancel(ctx context.Context, in *loms.OrderId) (*loms.EmptyResponse, error) {
	err := s.impl.OrderCancel(ctx, in.OrderId)
	if errors.Is(err, model.ErrorNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &loms.EmptyResponse{}, nil
}
