package initialization

import (
	"net/http"

	"github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/loms_service/loms_client"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/middleware"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/repository"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	product_client "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/client"
	product_service "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service"
)

func Build_server() *server.CartServer {
	validator := gody.NewValidator()
	err := validator.AddRules(rule.Min, rule.NotEmpty)
	if err != nil {
		panic(err)
	}

	cartRepo := repository.NewUserStorage()

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

	cartService := service.NewCartService(cartRepo, productService, wrappedLOMSClient)

	return server.NewCartServer(cartService, validator)
}
