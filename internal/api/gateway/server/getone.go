package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/pkg/api/gateway"
)

func (i *gatewayServer) MovieGetOne(ctx context.Context, req *pb.GatewayMovieGetOneRequest) (*pb.GatewayMovieGetOneResponse, error) {
	id := req.GetId()
	if id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id must be > 0")
	}

	m, err := i.storage.GetOneMovie(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GatewayMovieGetOneResponse{
		Movie: &pb.Movie{
			Id:    m.Id,
			Title: m.Title,
			Year: int32(m.Year),
		},
	}, nil
}
