package initialization

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
)

type IStockRepo interface {
	GetStock(context.Context, model.SKU) (model.Stock, error)
	InsertStock(context.Context, model.Stock) error
}

func Fill_stock_repo_from_json(ctx context.Context, r IStockRepo, path string) error {
	stocks, err := load_stock_json(path)
	if err != nil {
		return err
	}

	for _, stock := range stocks {
		err = r.InsertStock(ctx, stock)
		if err != nil {
			return err
		}
	}
	return nil
}

func load_stock_json(path string) ([]model.Stock, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var stocks []model.Stock

	err = json.Unmarshal(byteValue, &stocks)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}
