package initialization

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository_sqlc"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service"
)

func Build_server(ctx context.Context) *server.LOMSServer {
	time.Sleep(time.Second)
	conn, err := pgx.Connect(ctx, "postgres://user:password@pg_db:5432/loms_db")
	if err != nil {
		panic(err)
	}

	stockRepo := repository_sqlc.NewStockStorage(conn)
	orderRepo := repository_sqlc.NewOrderStorage(conn)

	stockS := service.NewStockService(stockRepo)
	orderS := service.NewOrderService(orderRepo)
	lomsService := loms_service.NewLOMSService(orderS, stockS)
	return server.NewLOMSServer(lomsService)
}
