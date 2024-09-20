package initialization

import (
	"context"
	"fmt"
	"path/filepath"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service"
)

func Build_server(ctx context.Context) *server.LOMSServer {
	stockRepo := repository.NewStockStorage()
	p, _ := filepath.Abs("stock-data.json")
	fmt.Println(p)
	err := Fill_stock_repo_from_json(ctx, stockRepo, p)
	if err != nil {
		panic(err)
	}

	orderRepo := repository.NewOrderStorage()
	stockS := service.NewStockService(stockRepo)
	orderS := service.NewOrderService(orderRepo)
	lomsService := loms_service.NewLOMSService(orderS, stockS)
	return server.NewLOMSServer(lomsService)
}
