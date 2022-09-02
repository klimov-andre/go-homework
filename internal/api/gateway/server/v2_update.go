package server

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/config/gateway"
	"homework/config/kafka"
	"homework/internal/api/gateway/metrics"
	pb "homework/pkg/api/gateway"
	pbStorage "homework/pkg/api/storage"

	"google.golang.org/protobuf/proto"
)

func (g *gatewayServer) MovieUpdateQueued(ctx context.Context, req *pb.GatewayMovieUpdateRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieUpdateQueued")
	defer span.End()

	metrics.GatewayTotalUpdateRequests.Add(1)

	m := req.GetMovie()
	if m.GetId() <= 0 {
		metrics.GatewayInvalidUpdateRequests.Add(1)
		return nil, status.Error(codes.InvalidArgument, "movie.id must be > 0")
	}

	request := &pbStorage.StorageMovieUpdateRequest{
		Movie: &pbStorage.Movie{
			Id:    m.GetId(),
			Title: m.GetTitle(),
			Year:  m.GetYear(),
		},
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		metrics.GatewayUnsuccessfulUpdateRequests.Add(1)
		span.RecordError(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	g.kafkaSender.SendMessage(ctx, kafka.TopicUpdate, fmt.Sprint(m.GetId()), msg)

	metrics.GatewaySuccessUpdateRequests.Add(1)
	return &emptypb.Empty{}, nil
}