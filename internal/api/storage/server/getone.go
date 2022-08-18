package server

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	storagePkg "homework/internal/storage"
	"homework/internal/storage/connections"
	pb "homework/pkg/api/storage"
)

func (i *storageServer) MovieGetOne(ctx context.Context, req *pb.StorageMovieGetOneRequest) (*pb.StorageMovieGetOneResponse, error) {
	m, err := i.storage.GetOneMovie(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, storagePkg.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	} else if errors.Is(err, connections.ErrTimeout) {
		return nil, status.Error(codes.DeadlineExceeded, err.Error())
	}

	return &pb.StorageMovieGetOneResponse{
		Movie: &pb.Movie{
			Id:    m.Id,
			Title: m.Title,
			Year: int32(m.Year),
		},
	}, nil
}
