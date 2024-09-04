package main

import (
	"log"
	"net/http"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/middleware"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/product"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/repository"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service"
)

// TODO: Добавить валидацию
func main() {
	log.Println("app starting")

	cartRepo := repository.NewUserStorage()

	client := http.Client{
		Transport: middleware.NewRetryRT(http.DefaultTransport),
		Timeout:   0,
	}

	productClient := product.NewProductClient(client, "http://route256.pavl.uk:8080", "testtoken")

	productService := service.NewProductService(productClient)
	cartService := service.NewCartService(cartRepo, productService)

	cartServer := server.NewCartServer(cartService)

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

// func test(cartService service.CartService) {
// 	cartService.AddItem(context.Background(), model.CartItem{
// 		SKU:    773297411,
// 		Count:  2,
// 		UserId: 1,
// 	})

// 	mas, _ := cartService.GetItems(context.Background(), 1)
// 	fmt.Println(mas)
// }
