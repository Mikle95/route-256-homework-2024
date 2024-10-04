package service

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/errorgroup"
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

	ch := make(chan model.ItemInfo)
	defer close(ch)

	eg := s.get_product_producer(ctx, ch, mas)
	read_ctx := s.get_product_consumer(ctx, ch, len(mas), &result)

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	<-read_ctx.Done()

	return &result, nil
}

func (s *CartService) get_product_producer(ctx context.Context, ch chan model.ItemInfo, mas []model.CartItem) *errorgroup.ErrorGroup {
	eg, eg_ctx := errorgroup.New(ctx)

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

	return eg
}

func (s *CartService) get_product_consumer(ctx context.Context, ch chan model.ItemInfo, count int, result *model.UserCartInfo) context.Context {
	read_ctx, cancel := context.WithCancel(ctx)
	go func() {
		for i := 0; i < count; i++ {
			item, ok := <-ch
			if !ok {
				break
			}
			result.Items = append(result.Items, item)
			result.TotalPrice += item.Price * uint32(item.Count)
		}
		cancel()
	}()
	return read_ctx
}
