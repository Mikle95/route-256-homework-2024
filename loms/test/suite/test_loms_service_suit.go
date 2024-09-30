package suite

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/model"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/initialization"
)

type LOMSServiceSuite struct {
	suite.Suite
	service *loms_service.LOMSService
}

func (s *LOMSServiceSuite) SetupTest() {
	ctx := context.Background()

	stockRepo := repository.NewStockStorage()

	p, _ := filepath.Abs("../../stock-data.json")
	fmt.Println(p)

	err := initialization.Fill_stock_repo_from_json(ctx, stockRepo, "../../stock-data.json")
	if err != nil {
		panic(err)
	}

	orderRepo := repository.NewOrderStorage()
	stockS := service.NewStockService(stockRepo)
	orderS := service.NewOrderService(orderRepo)

	s.service = loms_service.NewLOMSService(orderS, stockS)
}

func (s *LOMSServiceSuite) TestOrderCreate() {
	ctx := context.Background()

	expectedOrder := model.Order{
		User_id: 1,
		Items: []model.Item{
			{
				Sku:   1002,
				Count: 10,
			},
			{
				Sku:   1003,
				Count: 10,
			},
		},
	}

	id, err := s.service.OrderCreate(ctx, expectedOrder)
	require.NoError(s.T(), err)

	resultOrder, err := s.service.OrderInfo(ctx, id)
	require.NoError(s.T(), err)

	expectedOrder.Status = model.STATUS_WAIT
	assert.Equal(s.T(), expectedOrder, resultOrder)
}

func (s *LOMSServiceSuite) TestStockInfo() {
	ctx := context.Background()

	expectedOrder := model.Order{
		User_id: 1,
		Items: []model.Item{
			{
				Sku:   1002,
				Count: 10,
			},
		},
	}

	count, err := s.service.StocksInfo(ctx, 1002)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 180, int(count))

	_, err = s.service.OrderCreate(ctx, expectedOrder)
	require.NoError(s.T(), err)

	count, err = s.service.StocksInfo(ctx, 1002)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 170, int(count))
}

func (s *LOMSServiceSuite) TestOrderPay() {
	ctx := context.Background()

	expectedOrder := model.Order{
		User_id: 1,
		Items: []model.Item{
			{
				Sku:   1004,
				Count: 10,
			},
		},
	}

	count, err := s.service.StocksInfo(ctx, 1004)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 260, int(count))

	id, err := s.service.OrderCreate(ctx, expectedOrder)
	require.NoError(s.T(), err)

	err = s.service.OrderPay(ctx, id)
	require.NoError(s.T(), err)

	count, err = s.service.StocksInfo(ctx, 1004)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 250, int(count))
}

func (s *LOMSServiceSuite) TestOrdercancel() {
	ctx := context.Background()

	expectedOrder := model.Order{
		User_id: 1,
		Items: []model.Item{
			{
				Sku:   1004,
				Count: 10,
			},
		},
	}

	count, err := s.service.StocksInfo(ctx, 1004)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 260, int(count))

	id, err := s.service.OrderCreate(ctx, expectedOrder)
	require.NoError(s.T(), err)

	err = s.service.OrderCancel(ctx, id)
	require.NoError(s.T(), err)

	count, err = s.service.StocksInfo(ctx, 1004)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 260, int(count))
}
