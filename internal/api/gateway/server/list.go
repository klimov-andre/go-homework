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

func orderToString(order pb.ListOrder) string {
	result := "ASC"
	switch order {
	case pb.ListOrder_LIST_ORDER_DESC:
		result = "DESC"
	}

	return result
}

func (g *gatewayServer) MovieList(ctx context.Context, req *pb.GatewayMovieListRequest) (*pb.GatewayMovieListResponse, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieList")
	defer span.End()

	metrics.GatewayTotalListRequests.Add(1)

	limit := int(req.GetLimit())
	if limit <= 0 {
		return nil, status.Error(codes.InvalidArgument, "limit must be > 0")
	}

	order := orderToString(req.GetOrder())
	list, err := g.storage.List(ctx, limit, 0, order)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	result := make([]*pb.Movie, 0, len(list))
	for _, m := range list {
		result = append(result, &pb.Movie{
			Id:    m.Id,
			Title: m.Title,
			Year:  int32(m.Year),
		})
	}

	return &pb.GatewayMovieListResponse{
		Movie: result,
	}, nil
}