package server

import (
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	storageCfg "homework/config/storage"
	"homework/internal/storage/connections"
	pb "homework/pkg/api/storage"
)

func (s *storageServer) MovieList(req *pb.StorageMovieListRequest, srv pb.Storage_MovieListServer) error {
	ctx := srv.Context()

	var span trace.Span
	ctx, span = otel.Tracer(storageCfg.SpanTraceName).Start(ctx, "MovieList")
	defer span.End()

	limit, order := int(req.GetLimit()), req.GetOrder()
	for offset := int(req.GetOffset());; offset += limit {
		list, err := s.storage.List(ctx, limit, offset, order)
		if err != nil {
			span.RecordError(err)

			if errors.Is(err, connections.ErrTimeout) {
				return status.Error(codes.DeadlineExceeded, err.Error())
			}
			return status.Error(codes.Internal, err.Error())
		}

		// exit when no other movies in DB
		if len(list) == 0 {
			return nil
		}

		result := make([]*pb.Movie, 0, len(list))
		for _, m := range list {
			result = append(result, &pb.Movie{
				Id:    m.Id,
				Title: m.Title,
				Year:  int32(m.Year),
			})
		}

		srv.Send(&pb.StorageMovieListResponse{
			Movie: result,
		})
	}
}