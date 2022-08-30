package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
	pbStorage "homework/pkg/api/storage"

	"google.golang.org/protobuf/proto"
)

const topicCreate = "create"

func (g *gatewayServer) MovieCreateQueued(_ context.Context, req *pb.GatewayMovieCreateRequest) (*emptypb.Empty, error) {
	m, err := models.NewMovie(req.GetTitle(), int(req.GetYear()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	request := &pbStorage.StorageMovieCreateRequest{
		Title: m.Title,
		Year: int32(m.Year),
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	g.kafkaSender.SendMessage(topicCreate, "", msg)

	return &emptypb.Empty{}, nil
}