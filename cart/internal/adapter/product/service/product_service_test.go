package product

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service/mock"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

func TestProductService_GetItem(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	productRepoMock := mock.NewProductRepositoryMock(ctrl)

	sku := model.Sku(100)
	expectedItem := &domain.Item{SKU: 100, Name: "123", Price: 10}

	ps := NewProductService(productRepoMock)

	t.Run("Test valid get item", func(t *testing.T) {
		productRepoMock.GetProductMock.Expect(ctx, sku).Return(expectedItem, nil)
		item, err := ps.GetProduct(ctx, sku)

		require.NoError(t, err)
		assert.Equal(t, expectedItem, item)
	})

	t.Run("Test invalid sku get item", func(t *testing.T) {
		productRepoMock.GetProductMock.Expect(ctx, sku).Return(&domain.Item{}, nil)
		item, err := ps.GetProduct(ctx, sku)

		require.Equal(t, err.Error(), "sku does not exist")
		assert.Nil(t, item)
	})

	t.Run("Test invalid response get item", func(t *testing.T) {
		productRepoMock.GetProductMock.Expect(ctx, sku).Return(nil, errors.New(""))
		item, err := ps.GetProduct(ctx, sku)

		require.Empty(t, err)
		assert.Nil(t, item)
	})
}
