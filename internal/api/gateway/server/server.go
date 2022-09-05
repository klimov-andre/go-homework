package server

import (
	"homework/internal/api/gateway/kafka/sender"
	"homework/internal/storage/facade"
	pb "homework/pkg/api/gateway"
)

type gatewayServer struct {
	pb.UnimplementedGatewayServer

	storage     facade.StorageFacade
	kafkaSender *sender.Sender
}

func New(storage facade.StorageFacade, kafkaSender *sender.Sender) pb.GatewayServer {
	return &gatewayServer{
		storage: storage,
		kafkaSender: kafkaSender,
	}
}
