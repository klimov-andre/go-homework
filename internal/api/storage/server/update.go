package server

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	storageCfg "homework/config/storage"
	"homework/internal/storage"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (s *storageServer) MovieUpdate(ctx context.Context, req *pb.StorageMovieUpdateRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "MovieUpdate")
	defer span.End()

	m := req.GetMovie()
	upd := &models.Movie{
		Title: m.Title,
		Year: int(m.Year),
	}

	if err := s.storage.Update(ctx, m.GetId(), upd); err != nil {
		span.RecordError(err)

		if errors.Is(err, storage.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}