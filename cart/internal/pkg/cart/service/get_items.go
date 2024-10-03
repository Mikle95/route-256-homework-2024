package service

import (
	"context"
	"sync"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
	"golang.org/x/sync/errgroup"
)

func (s *CartService) GetItems(ctx context.Context, userId model.UID) (*model.UserCartInfo, error) {
	mas, err := s.repository.GetItems(ctx, userId)
	if err != nil {
		return nil, err
	}

	result := model.UserCartInfo{
		Items:      make([]model.ItemInfo, 0, len(mas)),
		TotalPrice: 0,
	}

	eg, eg_ctx := errgroup.WithContext(ctx)
	wg := sync.WaitGroup{}
	wg.Add(len(mas))

	ch := make(chan model.ItemInfo)
	defer close(ch)

	for _, val := range mas {
		eg.Go(func() error {
			item, err := s.productService.GetProduct(eg_ctx, val.SKU)
			if err != nil {
				return err
			}
			ch <- model.ItemInfo{
				SKU:   item.SKU,
				Name:  item.Name,
				Price: item.Price,
				Count: val.Count,
			}

			return nil
		})
	}

	eg.Go(func() error {
		for i := 0; i < len(mas); i++ {
			item, ok := <-ch
			if !ok {
				break
			}
			result.Items = append(result.Items, item)
			result.TotalPrice += item.Price * uint32(item.Count)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return &result, nil
}
