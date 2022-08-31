package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework/config/gateway"
	"homework/internal/api/gateway/kafka/sender"
	apiPkg "homework/internal/api/gateway/server"
	"homework/internal/storage/facade"
	"net"
	"net/http"

	pb "homework/pkg/api/gateway"
)

func runGRPC(storage facade.StorageFacade, kafkaSender *sender.Sender) {
	listener, err := net.Listen(gateway.GrpcProtocol, gateway.GrpcPort)
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGatewayServer(grpcServer, apiPkg.New(storage, kafkaSender))

	if err = grpcServer.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}

func runREST() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterGatewayHandlerFromEndpoint(ctx, mux, gateway.GrpcPort, opts); err != nil {
		logrus.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := http.ListenAndServe(gateway.RestPort, mux); err != nil {
		logrus.Fatal(err)
	}
}
