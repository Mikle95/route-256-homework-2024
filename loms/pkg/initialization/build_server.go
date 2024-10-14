package initialization

import (
	"context"

	"gitlab.ozon.dev/1mikle1/homework/loms/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service"
)

func Build_server(ctx context.Context) *server.LOMSServer {
	// pool, err := pgxpool.New(ctx, "postgres://user:password@pg_db:5432/loms_db")
	// if err != nil {
	// 	panic(err)
	// }

	// stockRepo := repository_sqlc.NewStockStorage(pool)
	// orderRepo := repository_sqlc.NewOrderStorage(pool)
	// TODO: Поменять обратно на бд
	stockRepo := repository.NewStockStorage()
	orderRepo := repository.NewOrderStorage()

	stockS := service.NewStockService(stockRepo)
	orderS := service.NewOrderService(orderRepo)
	lomsService := loms_service.NewLOMSService(orderS, stockS)
	return server.NewLOMSServer(lomsService)
}
