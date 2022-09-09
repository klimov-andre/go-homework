package server

import (
	"context"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	storageCfg "homework/config/storage"
	"homework/internal/storage"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (s *storageServer) MovieCreate(ctx context.Context, req *pb.StorageMovieCreateRequest) (*pb.StorageMovieCreateResponse, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "MovieCreate")
	defer span.End()

	m := &models.Movie{
		Title: req.GetTitle(),
		Year: int(req.GetYear()),
	}

	var (
		id uint64
		err error
	)
	if id, err = s.storage.Add(ctx, m); err != nil {
		span.RecordError(err)

		if errors.Is(err, storage.ErrMovieExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.publisher.Publish(m)
	return &pb.StorageMovieCreateResponse{Id: id}, nil
}