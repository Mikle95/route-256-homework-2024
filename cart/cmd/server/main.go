package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/http/middleware"
	"gitlab.ozon.dev/1mikle1/homework/cart/pkg/api/initialization"
)

func main() {
	log.Println("app starting")

	cartServer := initialization.Build_server()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItem)
	mux.HandleFunc("GET /user/{user_id}/cart", cartServer.GetItems)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartServer.DeleteCart)
	mux.HandleFunc("POST /cart/checkout", cartServer.Checkout)

	log.Println("server starting")

	logMux := middleware.NewLogMux(mux)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Println("app stopping")
	}()

	if err := http.ListenAndServe(":8082", logMux); err != nil {
		panic(err)
	}

}
