package service

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/mock"
)

func TestStockService_All(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	stockRepoMock := mock.NewIStockRepoMock(ctrl)
	stockService := NewStockService(stockRepoMock)

	stock := model.Stock{
		Sku:         1,
		Total_count: 15,
		Reserved:    2,
	}

	stockRepoMock.InsertStockMock.Expect(ctx, stock).Return(nil)
	stockRepoMock.GetStockMock.Expect(ctx, stock.Sku).Return(stock, nil)

	err := stockService.AddStock(ctx, stock)
	require.NoError(t, err)

	resultStock, err := stockService.GetStockBySKU(ctx, stock.Sku)
	require.NoError(t, err)
	require.Equal(t, stock, resultStock)

	items := []model.Item{
		{
			Sku:   stock.Sku,
			Count: 20,
		},
	}

	err = stockService.Reserve(ctx, items)
	require.Equal(t, "can't reserve more than total_count", err.Error())

	err = stockService.ReserveRemove(ctx, items)
	require.Equal(t, "can't remove more than reserved", err.Error())

	err = stockService.ReserveCancel(ctx, items)
	require.Equal(t, "can't cancel more than reserved", err.Error())

	items[0].Count = 5

	stock.Reserved = 7
	stockRepoMock.GetStockMock.Expect(ctx, stock.Sku).Return(stock, nil)

	stock.Reserved = 7 + 5
	stockRepoMock.InsertStockMock.Expect(ctx, stock).Return(nil)
	err = stockService.Reserve(ctx, items)
	require.NoError(t, err)

	stock.Reserved = 7 - 5
	stockRepoMock.InsertStockMock.Expect(ctx, stock).Return(nil)
	err = stockService.ReserveCancel(ctx, items)
	require.NoError(t, err)

	stock.Total_count = 15 - 5
	stock.Reserved = 7 - 5
	stockRepoMock.InsertStockMock.Expect(ctx, stock).Return(nil)
	err = stockService.ReserveRemove(ctx, items)
	require.NoError(t, err)
}
