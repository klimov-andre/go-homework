package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework/config"
	apiPkg "homework/internal/api/gateway/server"
	"homework/internal/storage/facade"
	"log"
	"net"
	"net/http"

	pb "homework/pkg/api/gateway"
)

func runGRPC(storage facade.StorageFacade) {
	listener, err := net.Listen(config.GrpcProtocol, config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGatewayServer(grpcServer, apiPkg.New(storage))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func runREST() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterGatewayHandlerFromEndpoint(ctx, mux, config.GrpcPort, opts); err != nil {
		log.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := http.ListenAndServe(config.RestPort, mux); err != nil {
		log.Fatal(err)
	}
}
