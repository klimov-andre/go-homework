package main

import (
	"google.golang.org/grpc"
	"homework/config"
	storagePkg "homework/internal/storage"
	"log"
	"net"

	apiPkg "homework/internal/api"
	pb "homework/pkg/api"
)

func runGRPC(storage *storagePkg.Storage) {
	listener, err := net.Listen(config.GrpcProtocol, config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(storage))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
