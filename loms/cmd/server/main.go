package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	middleware "gitlab.ozon.dev/1mikle1/homework/loms/internal/middlware"
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
	ctx := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.Logger,
			middleware.Validate,
		),
	)
	reflection.Register(grpcServer)

	server := initialization.Build_server(ctx)

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
		Handler: middleware.WithHTTPLoggingMiddleware(gwmux),
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)
	log.Fatalln(gwServer.ListenAndServe())
}
