package repository

import (
	"context"
	"testing"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

func BenchmarkCart_AddItem(b *testing.B) {
	ctx := context.Background()

	storage := NewCart()
	b.Run("Benchmark test AddItem", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			storage.AddItem(ctx, model.CartItem{})
		}
	})
}
