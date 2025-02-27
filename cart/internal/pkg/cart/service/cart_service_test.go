package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/loms_service/loms_client"
	product_client "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/client"
	product_service "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/middleware"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service/mock"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCartService_AddItem_table(t *testing.T) {
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
		{
			name: "wrong product service response",
			item: model.CartItem{
				SKU:    -1,
				UserId: 100,
				Count:  2,
			},
			expectedItem: model.CartItem{},
			wantErr:      errors.New("sku does not exist"),
		},
	}

	ctx := context.Background()
	ctrl := minimock.NewController(t)
	productServMock := mock.NewProductServiceMock(ctrl)
	cartRepoMock := mock.NewCartRepositoryMock(ctrl)
	lomsServiceMock := mock.NewILOMSServiceMock(ctrl)

	cs := NewCartService(cartRepoMock, productServMock, lomsServiceMock)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			productServMock.GetProductMock.Expect(ctx, tt.item.SKU).Return(&domain.Item{}, tt.wantErr)
			cartRepoMock.AddItemMock.Expect(ctx, tt.item).Return(tt.expectedItem, nil)
			lomsServiceMock.StocksInfoMock.Expect(ctx, tt.item.SKU).Return(100, nil)

			item, err := cs.AddItem(ctx, tt.item)

			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, item, tt.expectedItem)
		})
	}
}

func TestCartService_DeleteItem(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	productServMock := mock.NewProductServiceMock(ctrl)
	cartRepoMock := mock.NewCartRepositoryMock(ctrl)
	cs := NewCartService(cartRepoMock, productServMock, nil)

	t.Run("Delete item test", func(t *testing.T) {

		cartRepoMock.DeleteItemMock.Expect(ctx, 100, 100).Return(nil)
		err := cs.DeleteItem(ctx, 100, 100)
		require.NoError(t, err)
	})
}

func TestCartService_DeleteCart(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	productServMock := mock.NewProductServiceMock(ctrl)
	cartRepoMock := mock.NewCartRepositoryMock(ctrl)
	cs := NewCartService(cartRepoMock, productServMock, nil)

	t.Run("Delete cart test", func(t *testing.T) {

		cartRepoMock.DeleteCartMock.Expect(ctx, 100).Return(nil)
		err := cs.DeleteCart(ctx, 100)
		require.NoError(t, err)
	})
}

func TestCartService_GetItems(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	productServMock := mock.NewProductServiceMock(ctrl)
	cartRepoMock := mock.NewCartRepositoryMock(ctrl)
	cs := NewCartService(cartRepoMock, productServMock, nil)

	userId := int64(100)

	item := model.CartItem{
		SKU:    100,
		UserId: 100,
		Count:  10,
	}

	itemInfo := &model.UserCartInfo{
		Items: []model.ItemInfo{
			{
				SKU:   100,
				Name:  "Bottle",
				Price: 10,
				Count: 10,
			},
		},
		TotalPrice: 100,
	}

	t.Run("GetItems test", func(t *testing.T) {

		cartRepoMock.GetItemsMock.Expect(ctx, userId).Return([]model.CartItem{
			item,
		}, nil)
		productServMock.GetProductMock.Expect(ctx, item.SKU).Return(&domain.Item{
			SKU: itemInfo.Items[0].SKU, Price: itemInfo.Items[0].Price, Name: itemInfo.Items[0].Name}, nil)

		items, err := cs.GetItems(ctx, userId)

		require.NoError(t, err)
		assert.Equal(t, itemInfo, items)

		productServMock.GetProductMock.Expect(ctx, item.SKU).Return(nil, errors.New(""))
		items, err = cs.GetItems(ctx, userId)

		require.Error(t, err)
		assert.Nil(t, items)

		cartRepoMock.GetItemsMock.Expect(ctx, userId).Return(nil, errors.New(""))
		items, err = cs.GetItems(ctx, userId)

		require.Error(t, err)
		assert.Nil(t, items)

	})
}

func TestCartService_GetItems_parallel(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)
	cartRepoMock := mock.NewCartRepositoryMock(ctrl)

	client := http.Client{
		Transport: middleware.NewRetryRT(http.DefaultTransport),
		Timeout:   0,
	}

	productClient := product_client.NewProductClient(client, "http://route256.pavl.uk:8080", "testtoken", 10)
	productService := product_service.NewProductService(productClient)

	// gRPC client
	conn, err := grpc.NewClient("loms:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	lomsClient := loms.NewLOMSClient(conn)
	wrappedLOMSClient := loms_client.NewClient("user", lomsClient)
	cartService := NewCartService(cartRepoMock, productService, wrappedLOMSClient)

	items := make([]model.CartItem, 10)
	for i := range 8 {
		items[i] = model.CartItem{SKU: 773297411}
	}

	cartRepoMock.GetItemsMock.Expect(ctx, 0).Return(items, nil)
	_, err = cartService.GetItems(ctx, 0)
	require.Error(t, err)
	fmt.Printf("%s\n", err.Error())
}
