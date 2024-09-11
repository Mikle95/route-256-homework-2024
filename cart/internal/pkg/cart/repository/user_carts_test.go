package repository

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

func TestHandler_AddItem_Table(t *testing.T) {
	ctx := context.Background()

	type data struct {
		name          string
		item          model.CartItem
		wantErr       error
		expectedCount model.Count
	}

	testData := []data{
		{
			name: "valid add new item",
			item: model.CartItem{
				UserId: 123,
				SKU:    100,
				Count:  2,
			},
			wantErr:       nil,
			expectedCount: 2,
		},
		{
			name: "valid add existing item",
			item: model.CartItem{
				UserId: 123,
				SKU:    100,
				Count:  2,
			},
			wantErr:       nil,
			expectedCount: 4,
		},
	}

	userStorage := NewUserStorage()

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			cItem, err := userStorage.AddItem(ctx, tt.item)

			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assert.Equal(t, cItem.Count, tt.expectedCount)
			}
		})
	}
}

func TestUserCart_DeleteItem_Table(t *testing.T) {
	ctx := context.Background()

	type data struct {
		name          string
		UserId        model.UID
		SKU           model.Sku
		expectedCount model.Count
	}

	testData := []data{
		{
			name:          "delete cart item (in cart)",
			UserId:        123,
			SKU:           100,
			expectedCount: 2,
		},
		{
			name:          "delete item (not in cart)",
			UserId:        123,
			SKU:           100,
			expectedCount: 2,
		},
	}

	userStorage := NewUserStorage()
	userStorage.AddItem(ctx, model.CartItem{
		SKU:    100,
		UserId: 123,
		Count:  2,
	})

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := userStorage.DeleteItem(ctx, tt.UserId, tt.SKU)
			require.NoError(t, err)
		})
	}
}

func TestUserCart_GetItems(t *testing.T) {
	ctx := context.Background()
	userStorage := NewUserStorage()

	userId := model.Sku(123)

	items := []model.CartItem{
		{
			SKU:    100,
			UserId: userId,
			Count:  3,
		},
		{
			SKU:    101,
			UserId: userId,
			Count:  5,
		},
		{
			SKU:    100,
			UserId: userId,
			Count:  6,
		},
	}

	expectedItems := []model.CartItem{
		{
			SKU:    100,
			UserId: userId,
			Count:  9,
		},
		{
			SKU:    101,
			UserId: userId,
			Count:  5,
		},
	}

	t.Run("Get items", func(t *testing.T) {

		for _, item := range items {
			userStorage.AddItem(ctx, item)
		}
		resItems, err := userStorage.GetItems(ctx, 123)
		require.NoError(t, err)

		sort.Slice(expectedItems, func(i, j int) bool {
			return expectedItems[i].SKU < expectedItems[j].SKU
		})

		sort.Slice(resItems, func(i, j int) bool {
			return resItems[i].SKU < resItems[j].SKU
		})

		assert.Equal(t, expectedItems, resItems)
	})
}

func TestUserCart_DeleteCart_Table(t *testing.T) {
	ctx := context.Background()

	type data struct {
		name          string
		UserId        model.UID
		expectedCount model.Count
	}

	testData := []data{
		{
			name:          "delete cart (exist)",
			UserId:        123,
			expectedCount: 2,
		},
		{
			name:          "delete cart (not exist)",
			UserId:        123,
			expectedCount: 2,
		},
	}

	userStorage := NewUserStorage()
	userStorage.AddItem(ctx, model.CartItem{
		SKU:    100,
		UserId: 123,
		Count:  2,
	})

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := userStorage.DeleteCart(ctx, tt.UserId)
			require.NoError(t, err)
			mas, err := userStorage.GetItems(ctx, tt.UserId)
			require.NoError(t, err)
			assert.Empty(t, mas)
		})
	}
}
