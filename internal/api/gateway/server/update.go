package server

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/config/gateway"
	"homework/internal/api/gateway/metrics"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieUpdate(ctx context.Context, req *pb.GatewayMovieUpdateRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieUpdate")
	defer span.End()

	metrics.GatewayTotalUpdateRequests.Add(1)

	m := req.GetMovie()
	if m.GetId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "movie.id must be > 0")
	}

	// NewMovie check input params
	upd, err := models.NewMovie(m.GetTitle(), int(m.GetYear()))
	if err != nil {
		span.RecordError(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = g.storage.Update(ctx, m.GetId(), upd); err != nil {
		span.RecordError(err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}