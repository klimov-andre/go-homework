package server

import (
	"homework/internal/storage/facade"
	pb "homework/pkg/api/gateway"
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
