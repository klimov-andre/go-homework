package api

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	storagePkg "homework/internal/storage"
	"homework/internal/storage/connections"
	pb "homework/pkg/api"
)

func (i *implementation) MovieCreate(_ context.Context, req *pb.MovieCreateRequest) (*pb.MovieCreateResponse, error) {
	if _, err := i.storage.Add(req.GetTitle(), int(req.GetYear())); err != nil {
		if errors.Is(err, storagePkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, storagePkg.ErrMovieExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MovieCreateResponse{}, nil
}