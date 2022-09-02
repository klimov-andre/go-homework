package server

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/config/gateway"
	"homework/internal/api/gateway/metrics"
	pb "homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieGetOne(ctx context.Context, req *pb.GatewayMovieGetOneRequest) (*pb.GatewayMovieGetOneResponse, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieGetOne")
	defer span.End()

	metrics.GatewayTotalGetOneRequests.Add(1)

	id := req.GetId()
	if id <= 0 {
		metrics.GatewayInvalidGetOneRequests.Add(1)
		return nil, status.Error(codes.InvalidArgument, "id must be > 0")
	}

	m, err := g.storage.GetOneMovie(ctx, req.GetId())
	if err != nil {
		metrics.GatewayUnsuccessfulGetOneRequests.Add(1)
		span.RecordError(err)
		return nil, err
	}

	metrics.GatewaySuccessGetOneRequests.Add(1)
	return &pb.GatewayMovieGetOneResponse{
		Movie: &pb.Movie{
			Id:    m.Id,
			Title: m.Title,
			Year: int32(m.Year),
		},
	}, nil
}
