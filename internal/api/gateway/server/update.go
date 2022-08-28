package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieUpdate(ctx context.Context, req *pb.GatewayMovieUpdateRequest) (*emptypb.Empty, error) {
	m := req.GetMovie()
	if m.GetId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "movie.id must be > 0")
	}

	// NewMovie check input params
	upd, err := models.NewMovie(m.GetTitle(), int(m.GetYear()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = g.storage.Update(ctx, m.GetId(), upd); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}