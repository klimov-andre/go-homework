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

func (g *gatewayServer) MovieUpdateQueued(_ context.Context, req *pb.GatewayMovieUpdateRequest) (*emptypb.Empty, error) {
	m := req.GetMovie()
	request := &pbStorage.StorageMovieUpdateRequest{
		Movie: &pbStorage.Movie{
			Id:    m.GetId(),
			Title: m.GetTitle(),
			Year:  m.GetYear(),
		},
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	g.kafkaSender.SendMessage(kafka.TopicUpdate, fmt.Sprint(m.GetId()), msg)
	return &emptypb.Empty{}, nil
}