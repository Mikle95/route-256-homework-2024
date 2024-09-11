package service

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service/mock"
)

func TestCartService_AddItem(t *testing.T) {
	type data struct {
		name         string
		item         model.CartItem
		expectedItem model.CartItem
		wantErr      error
	}

	testData := []data{
		{
			name: "add valid item",
			item: model.CartItem{
				SKU:    100,
				UserId: 100,
				Count:  2,
			},
			expectedItem: model.CartItem{
				SKU:    100,
				UserId: 100,
				Count:  2,
			},
		},
	}

	ctx := context.Background()
	ctrl := minimock.NewController(t)
	productServMock := mock.NewProductServiceMock(ctrl)
	cartRepoMock := mock.NewCartRepositoryMock(ctrl)
	cs := NewCartService(cartRepoMock, productServMock)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			productServMock.GetProductMock.Expect(ctx, tt.item.SKU).Return(&domain.Item{}, tt.wantErr)
			cartRepoMock.AddItemMock.Expect(ctx, tt.item).Return(tt.expectedItem, nil)

			item, err := cs.AddItem(ctx, tt.item)

			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, item, tt.expectedItem)
		})
	}
}
