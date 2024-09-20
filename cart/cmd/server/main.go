package main

import (
	"log"
	"net/http"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/loms_service/loms_client"
	product_client "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/client"
	product_service "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/middleware"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/repository"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gody "github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
)

func main() {
	log.Println("app starting")

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

	productClient := product_client.NewProductClient(client, "http://route256.pavl.uk:8080", "testtoken")

	productService := product_service.NewProductService(productClient)

	// gRPC client
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	lomsClient := loms.NewLOMSClient(conn)
	wrappedLOMSClient := loms_client.NewClient("user", lomsClient)

	cartService := service.NewCartService(cartRepo, productService, wrappedLOMSClient)

	cartServer := server.NewCartServer(cartService, validator)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItem)
	mux.HandleFunc("GET /user/{user_id}/cart", cartServer.GetItems)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartServer.DeleteCart)
	mux.HandleFunc("POST /cart/checkout", cartServer.Checkout)

	log.Println("server starting")

	logMux := middleware.NewLogMux(mux)

	if err := http.ListenAndServe(":8082", logMux); err != nil {
		panic(err)
	}

}
