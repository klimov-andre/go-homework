package server

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/internal/storage"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (i *storageServer) MovieUpdate(ctx context.Context, req *pb.StorageMovieUpdateRequest) (*emptypb.Empty, error) {
	m := req.GetMovie()
	upd := &models.Movie{
		Title: m.Title,
		Year: int(m.Year),
	}

	if err := i.storage.Update(ctx, m.GetId(), upd); err != nil {
		if errors.Is(err, storage.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}