package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/app/server"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/mw"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/repository"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service"
	"gitlab.ozon.dev/1mikle1/homework/loms/internal/service/loms_service"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/api/loms/v1"
	"gitlab.ozon.dev/1mikle1/homework/loms/pkg/initialization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50051
	httpPort = 8081
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// mw.Panic,
			mw.Logger,
			mw.Validate,
		),
	)
	reflection.Register(grpcServer)

	ctx := context.Background()

	stockRepo := repository.NewStockStorage()
	p, _ := filepath.Abs("stock-data.json")
	fmt.Println(p)
	err = initialization.Fill_stock_repo_from_json(ctx, stockRepo, p)
	if err != nil {
		panic(err)
	}

	orderRepo := repository.NewOrderStorage()
	stockS := service.NewStockService(stockRepo)
	orderS := service.NewOrderService(orderRepo)
	lomsService := loms_service.NewLOMSService(orderS, stockS)
	server := server.NewLOMSServer(lomsService)
	fmt.Print(server)

	loms.RegisterLOMSServer(grpcServer, server)

	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to deal:", err)
	}

	gwmux := runtime.NewServeMux()

	if err = loms.RegisterLOMSHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: mw.WithHTTPLoggingMiddleware(gwmux),
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)
	log.Fatalln(gwServer.ListenAndServe())
}
