package loms_service

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service/mock"
)

func TestLOMSService_All(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	orderServiceMock := mock.NewIOrderServiceMock(ctrl)
	stockServiceMock := mock.NewIStockServiceMock(ctrl)

	service := NewLOMSService(orderServiceMock, stockServiceMock)

	sku := uint32(1002)

	expected_order := model.Order{User_id: 1, Items: []model.Item{
		{
			Sku:   sku,
			Count: 10,
		},
	}}

	expected_id := model.OID(1)

	orderServiceMock.CreateMock.Expect(ctx, expected_order).Return(expected_id, nil)
	stockServiceMock.ReserveMock.Expect(ctx, expected_order.Items).Return(nil)
	orderServiceMock.SetStatusMock.Expect(ctx, expected_id, model.STATUS_WAIT).Return(nil)
	id, err := service.OrderCreate(ctx, expected_order)
	require.NoError(t, err)
	require.EqualValues(t, expected_id, id)

	expected_order.Status = model.STATUS_WAIT
	orderServiceMock.GetByIdMock.Expect(ctx, expected_id).Return(expected_order, nil)
	order, err := service.OrderInfo(ctx, expected_id)
	require.NoError(t, err)
	require.Equal(t, expected_order, order)

	stockServiceMock.ReserveRemoveMock.Expect(ctx, expected_order.Items).Return(nil)
	orderServiceMock.SetStatusMock.Expect(ctx, expected_id, model.STATUS_PAYED).Return(nil)
	err = service.OrderPay(ctx, expected_id)
	require.NoError(t, err)

	stockServiceMock.ReserveCancelMock.Expect(ctx, expected_order.Items).Return(nil)
	orderServiceMock.SetStatusMock.Expect(ctx, expected_id, model.STATUS_CANCELLED).Return(nil)
	err = service.OrderCancel(ctx, expected_id)
	require.NoError(t, err)

	stock := model.Stock{Sku: sku, Total_count: 10, Reserved: 2}
	stockServiceMock.GetStockBySKUMock.Expect(ctx, sku).Return(stock, nil)
	count, err := service.StocksInfo(ctx, sku)
	require.NoError(t, err)
	require.Equal(t, stock.Total_count-stock.Reserved, count)
}
