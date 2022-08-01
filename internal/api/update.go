package api

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	storagePkg "homework/internal/storage"
	"homework/internal/storage/connections"
	pb "homework/pkg/api"
)

func (i *implementation) MovieUpdate(_ context.Context, req *pb.MovieUpdateRequest) (*pb.MovieUpdateResponse, error) {
	m := req.GetMovie()
	if _, err := i.storage.Update(m.GetId(), m.GetTitle(), int(m.GetYear())); err != nil {
		if errors.Is(err, storagePkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, storagePkg.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MovieUpdateResponse{}, nil
}