package api

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	pb "homework/pkg/api"
)

func (i *implementation) MovieUpdate(_ context.Context, req *pb.MovieUpdateRequest) (*pb.MovieUpdateResponse, error) {
	m := req.GetMovie()

	upd, err := models.NewMovie(m.GetTitle(), int(m.GetYear()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := i.storage.Update(m.GetId(), upd); err != nil {
		if errors.Is(err, storage.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MovieUpdateResponse{}, nil
}