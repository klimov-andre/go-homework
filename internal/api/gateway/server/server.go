package server

import (
	"homework/internal/storage/facade"
	pb "homework/pkg/api"
)

type gatewayServer struct {
	pb.UnimplementedGatewayServer

	storage facade.StorageFacade
}

func New(storage facade.StorageFacade) pb.GatewayServer {
	return &gatewayServer{
		storage: storage,
	}
}
