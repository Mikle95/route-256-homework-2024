package suite

import (
	"context"
	"sort"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

func (s *ItemServiceSuite) TestGetItems() {
	ctx := context.Background()
	userID := int64(123)

	itemList, err := s.service.GetItems(ctx, userID)

	require.NoError(s.T(), err)

	sort.Slice(itemList.Items, func(i, j int) bool {
		return itemList.Items[i].SKU < itemList.Items[j].SKU
	})

	expectedItems := &model.UserCartInfo{
		Items: []model.ItemInfo{
			{
				SKU:   1148162,
				Name:  "Кулинар Гуров",
				Price: 2931,
				Count: 1,
			},
			{
				SKU:   773297411,
				Name:  "Кроссовки Nike JORDAN",
				Price: 2202,
				Count: 2,
			},
		},
		TotalPrice: 7335,
	}

	require.Equal(s.T(), expectedItems, itemList)
}
