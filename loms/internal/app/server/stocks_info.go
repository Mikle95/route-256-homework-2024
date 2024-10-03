package server

import (
	"context"
	"errors"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LOMSServer) StocksInfo(ctx context.Context, in *loms.StockIdRequest) (*loms.StocksInfoResponse, error) {
	stocks, err := s.impl.StocksInfo(ctx, in.Sku)
	if errors.Is(err, model.ErrorNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &loms.StocksInfoResponse{Count: stocks}, nil
}
