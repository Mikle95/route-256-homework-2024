package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

var host_name string = "http://route256.pavl.uk:8080"

type RequestData struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku"`
}

type ResponseData struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type ProductClient struct {
	hostName string
	token    string
}

func NewProductClient(host string, token string) *ProductClient {
	return &ProductClient{hostName: host, token: token}
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

	resp, err := http.Post(host_name+path_api, "application/json", bytes.NewBuffer(jsonBody))
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
