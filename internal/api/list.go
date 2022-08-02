package api

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage/connections"
	pb "homework/pkg/api"
)

func (i *implementation) MovieList(_ context.Context, _ *pb.MovieListRequest) (*pb.MovieListResponse, error) {
	list, err := i.storage.List()
	if err != nil {
		if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.Movie, 0, len(list))
	for _, m := range list {
		result = append(result, &pb.Movie{
			Id:    m.Id(),
			Title: m.Title(),
			Year:  int32(m.Year()),
		})
	}

	return &pb.MovieListResponse{
		Movie: result,
	}, nil
}