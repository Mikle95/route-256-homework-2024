package loms_client

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc/metadata"
)

func (c *Client) StocksInfo(ctx context.Context, sku domain.Sku) (uint64, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-auth", c.header)
	response, err := c.client.StocksInfo(ctx, &loms.StockIDRequest{Info: &loms.StockID{Sku: uint32(sku)}})
	if err != nil {
		return 0, err
	}
	return response.Count, nil
}
