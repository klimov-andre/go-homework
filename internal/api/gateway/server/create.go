package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
)

func (i *gatewayServer) MovieCreate(ctx context.Context, req *pb.GatewayMovieCreateRequest) (*pb.GatewayMovieCreateResponse, error) {
	// NewMovie method checks input params
	m, err := models.NewMovie(req.GetTitle(), int(req.GetYear()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var id uint64
	if id, err = i.storage.Add(ctx, m); err != nil {
		return nil, err
	}
	return &pb.GatewayMovieCreateResponse{
		Id: id,
	}, nil
}