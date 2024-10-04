package product

import (
	"bytes"
	"context"
	"encoding/json"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
)

type RequestData struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku"`
}

type ResponseData struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

const path_get_product_api = "/get_product"

func (p *ProductClient) GetProduct(ctx context.Context, sku domain.Sku) (rd *domain.Item, err error) {

	if err := p.limiter.Wait(ctx); err != nil {
		return nil, err
	}

	body := RequestData{
		Token: p.token,
		Sku:   int64(sku),
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Post(p.hostName+path_get_product_api, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}

	result := domain.Item{
		SKU:   sku,
		Name:  responseData.Name,
		Price: responseData.Price,
	}

	return &result, nil
}
