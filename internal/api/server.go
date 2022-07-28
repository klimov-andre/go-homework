package api

import (
	"context"
	pb "homework/pkg/api"
)

type implementation struct {
	pb.UnimplementedAdminServer
}

func New() pb.AdminServer {
	return &implementation{}
}

func (i *implementation) MovieCreate(_ context.Context, _ *pb.MovieCreateRequest) (*pb.MovieCreateResponse, error) {
	return &pb.MovieCreateResponse{}, nil
}

func (i *implementation) MovieList(_ context.Context, _ *pb.MovieListRequest) (*pb.MovieListResponse, error) {
	return &pb.MovieListResponse{}, nil
}

func (i *implementation) MovieUpdate(_ context.Context, _ *pb.MovieUpdateRequest) (*pb.MovieUpdateResponse, error) {
	return &pb.MovieUpdateResponse{}, nil
}

func (i *implementation) MovieDelete(_ context.Context, _ *pb.MovieDeleteRequest) (*pb.MovieDeleteResponse, error) {
	return &pb.MovieDeleteResponse{}, nil
}
