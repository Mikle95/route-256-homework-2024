package loms_client

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc/metadata"
)

func (c *Client) Checkout(ctx context.Context, order domain.Order) (domain.OID, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "x-auth", c.header)
	response, err := c.client.OrderCreate(ctx, orderRepack(order))
	if err != nil {
		return 0, err
	}
	return response.OrderID, err
}

func orderRepack(order domain.Order) *loms.OrderInfoMessage {
	result := loms.OrderInfoMessage{User: order.User_id, Items: make([]*loms.Item, 0)}

	for _, item := range order.Items {
		result.Items = append(result.Items, &loms.Item{
			Sku:   uint32(item.Sku),
			Count: uint32(item.Count),
		})
	}

	return &result
}
