package server

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (i *storageServer) MovieCreate(ctx context.Context, req *pb.StorageMovieCreateRequest) (*pb.StorageMovieCreateResponse, error) {
	m := &models.Movie{
		Title: req.GetTitle(),
		Year: int(req.GetYear()),
	}

	var (
		id uint64
		err error
	)
	if id, err = i.storage.Add(ctx, m); err != nil {
		if errors.Is(err, storage.ErrMovieExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.StorageMovieCreateResponse{Id: id}, nil
}