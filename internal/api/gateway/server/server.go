package server

import (
	"homework/internal/api/gateway/kafka/sender"
	"homework/internal/api/gateway/subscriber"
	"homework/internal/storage/facade"
	pb "homework/pkg/api/gateway"
)

type gatewayServer struct {
	pb.UnimplementedGatewayServer

	storage     facade.StorageFacade
	kafkaSender *sender.Sender
	subscriber  *subscriber.Subscriber
}

func New(storage facade.StorageFacade, kafkaSender *sender.Sender) pb.GatewayServer {
	return &gatewayServer{
		storage: storage,
		kafkaSender: kafkaSender,
		subscriber: subscriber.NewRedisSubscriber(),
	}
}
