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
	pb "homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieDelete(ctx context.Context, req *pb.GatewayMovieDeleteRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieDelete")
	defer span.End()

	metrics.GatewayTotalDeleteRequests.Add(1)

	id := req.GetId()
	if id <= 0 {
		metrics.GatewayInvalidDeleteRequests.Add(1)
		return nil, status.Error(codes.InvalidArgument, "id must be > 0")
	}

	if err := g.storage.Delete(ctx, req.GetId()); err != nil {
		metrics.GatewayUnsuccessfulDeleteRequests.Add(1)
		span.RecordError(err)
		return nil, err
	}

	metrics.GatewaySuccessDeleteRequests.Add(1)
	return &emptypb.Empty{}, nil
}
