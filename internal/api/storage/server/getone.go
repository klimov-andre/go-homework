package server

import (
	"context"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	storageCfg "homework/config/storage"
	storagePkg "homework/internal/storage"
	"homework/internal/storage/connections"
	pb "homework/pkg/api/storage"
)

func (s *storageServer) MovieGetOne(ctx context.Context, req *pb.StorageMovieGetOneRequest) (*pb.StorageMovieGetOneResponse, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "MovieGetOne")
	defer span.End()

	m, err := s.storage.GetOneMovie(ctx, req.GetId())
	if err != nil {
		span.RecordError(err)

		if errors.Is(err, storagePkg.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.StorageMovieGetOneResponse{
		Movie: &pb.Movie{
			Id:    m.Id,
			Title: m.Title,
			Year: int32(m.Year),
		},
	}, nil
}
