package loms_client

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (c *Client) StocksInfo(ctx context.Context, sku domain.Sku) (uint64, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-auth", c.header)
	response, err := c.client.StocksInfo(ctx, &loms.StockIdRequest{Sku: uint32(sku)})
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.FailedPrecondition {
			return 0, fmt.Errorf("stocks info, %w: %w", domain.ErrorPrecondition, err)
		}

		return 0, err
	}
	return response.Count, nil
}
