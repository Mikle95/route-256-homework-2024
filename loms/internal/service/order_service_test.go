package service

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/mock"
)

func TestOrderService_All(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	orderRepoMock := mock.NewIOrderRepoMock(ctrl)
	orderService := NewOrderService(orderRepoMock)

	order := model.Order{User_id: 1, Items: []model.Item{
		{
			Sku:   1002,
			Count: 10,
		},
	}}

	expectOrder := order
	expectOrder.Status = model.STATUS_NEW

	orderRepoMock.AddOrderMock.Expect(ctx, expectOrder).Return(1, nil)
	id, err := orderService.Create(ctx, order)

	require.NoError(t, err)
	require.EqualValues(t, 1, id)

	expectOrder.Status = model.STATUS_WAIT
	orderRepoMock.GetOrderMock.Expect(ctx, id).Return(expectOrder, nil)
	order, err = orderService.GetById(ctx, id)
	require.NoError(t, err)
	require.Equal(t, expectOrder, order)

	expectOrder.Status = model.STATUS_CANCELLED
	orderRepoMock.ChangeOrderMock.Expect(ctx, id, expectOrder).Return(nil)
	err = orderService.SetStatus(ctx, id, model.STATUS_CANCELLED)
	require.NoError(t, err)
}
