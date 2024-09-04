package product

import (
	"bytes"
	"context"
	"encoding/json"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type RequestData struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku"`
}

type ResponseData struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (p *ProductClient) GetProduct(ctx context.Context, sku model.Sku) (rd *model.Item, err error) {
	path_api := "/get_product"

	body := RequestData{
		Token: p.token,
		Sku:   int64(sku),
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Post(p.hostName+path_api, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}

	result := model.Item{
		SKU:   sku,
		Name:  responseData.Name,
		Price: responseData.Price,
	}

	return &result, nil
}
