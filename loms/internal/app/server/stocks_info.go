package server

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LOMSServer) StocksInfo(ctx context.Context, in *loms.StockIdRequest) (*loms.StocksInfoResponse, error) {
	stocks, err := s.impl.StocksInfo(ctx, in.Sku)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	return &loms.StocksInfoResponse{Count: stocks}, nil
}
