package loms_client

import "gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"

type Client struct {
	header string
	client loms.LOMSClient
}

func NewClient(header string, client loms.LOMSClient) *Client {
	return &Client{client: client, header: header}
}
