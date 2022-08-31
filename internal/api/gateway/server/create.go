package server

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/config/gateway"
	"homework/internal/api/gateway/metrics"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieCreate(ctx context.Context, req *pb.GatewayMovieCreateRequest) (*pb.GatewayMovieCreateResponse, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieCreate")
	defer span.End()

	metrics.GatewayTotalAddRequests.Add(1)

	// NewMovie method checks input params
	m, err := models.NewMovie(req.GetTitle(), int(req.GetYear()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var id uint64
	if id, err = g.storage.Add(ctx, m); err != nil {
		span.RecordError(err)
		return nil, err
	}

	return &pb.GatewayMovieCreateResponse{
		Id: id,
	}, nil
}