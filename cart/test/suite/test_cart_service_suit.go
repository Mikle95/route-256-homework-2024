package suite

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	productC "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/client"
	productS "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/repository"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service"
)

type ItemServiceSuite struct {
	suite.Suite
	service *service.CartService
}

func (s *ItemServiceSuite) SetupSuite() {
	storage := repository.NewUserStorage()
	productClient := productC.NewProductClient(http.Client{}, "http://route256.pavl.uk:8080", "testtoken")
	productService := productS.NewProductService(productClient)
	s.service = service.NewCartService(storage, productService, nil)

	ctx := context.Background()

	userID := int64(123)
	item1 := model.CartItem{
		SKU:    773297411,
		Count:  2,
		UserId: userID,
	}

	_, err := s.service.AddItem(ctx, item1)
	require.NoError(s.T(), err)
	itemList, _ := s.service.GetItems(ctx, userID)

	require.Equal(s.T(), len(itemList.Items), 1)
	require.Equal(s.T(), itemList.Items[0].SKU, item1.SKU)
	require.Equal(s.T(), itemList.Items[0].Count, item1.Count)

	item2 := model.CartItem{
		SKU:    1148162,
		Count:  1,
		UserId: userID,
	}

	_, err = s.service.AddItem(ctx, item2)

	require.NoError(s.T(), err)
}
