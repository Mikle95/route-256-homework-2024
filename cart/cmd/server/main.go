package main

import (
	"log"
	"net/http"

	product_client "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/client"
	product_service "gitlab.ozon.dev/1mikle1/homework/cart/internal/adapter/product/service"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/middleware"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/repository"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service"

	gody "github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
)

func main() {
	log.Println("app starting")

	validator := gody.NewValidator()
	validator.AddRules(rule.Min, rule.NotEmpty)

	cartRepo := repository.NewUserStorage()

	client := http.Client{
		Transport: middleware.NewRetryRT(http.DefaultTransport),
		Timeout:   0,
	}

	productClient := product_client.NewProductClient(client, "http://route256.pavl.uk:8080", "testtoken")

	productService := product_service.NewProductService(productClient)
	cartService := service.NewCartService(cartRepo, productService)

	cartServer := server.NewCartServer(cartService, validator)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItem)
	mux.HandleFunc("GET /user/{user_id}/cart", cartServer.GetItems)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartServer.DeleteCart)

	log.Println("server starting")

	logMux := middleware.NewLogMux(mux)

	if err := http.ListenAndServe(":8082", logMux); err != nil {
		panic(err)
	}

}
