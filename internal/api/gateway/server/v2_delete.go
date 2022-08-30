package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/config/kafka"
	pb "homework/pkg/api/gateway"
	pbStorage "homework/pkg/api/storage"

	"google.golang.org/protobuf/proto"
)

func (g *gatewayServer) MovieDeleteQueued(_ context.Context, req *pb.GatewayMovieDeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	request := &pbStorage.StorageMovieDeleteRequest{
		Id: id,
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	g.kafkaSender.SendMessage(kafka.TopicDelete, fmt.Sprint(id), msg)
	return &emptypb.Empty{}, nil
}