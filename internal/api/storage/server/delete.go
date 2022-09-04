package server

import (
	"context"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	storageCfg "homework/config/storage"
	"homework/internal/storage/connections"
	pb "homework/pkg/api/storage"
)

func (s *storageServer) MovieDelete(ctx context.Context, req *pb.StorageMovieDeleteRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "MovieDelete")
	defer span.End()

	if err := s.storage.Delete(ctx, req.GetId()); err != nil {
		span.RecordError(err)

		if errors.Is(err, connections.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
