package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/config/gateway"
	"homework/config/kafka"
	"homework/internal/api/gateway/metrics"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
	pbStorage "homework/pkg/api/storage"

	"google.golang.org/protobuf/proto"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func (g *gatewayServer) MovieCreateQueued(ctx context.Context, req *pb.GatewayMovieCreateRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieCreateQueued")
	defer span.End()

	metrics.GatewayTotalAddRequests.Add(1)

	m, err := models.NewMovie(req.GetTitle(), int(req.GetYear()))
	if err != nil {
		span.RecordError(err)
		metrics.GatewayInvalidAddRequests.Add(1)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	request := &pbStorage.StorageMovieCreateRequest{
		Title: m.Title,
		Year: int32(m.Year),
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		metrics.GatewayUnsuccessfulAddRequests.Add(1)
		span.RecordError(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	g.kafkaSender.SendMessage(ctx, kafka.TopicCreate, "", msg)

	metrics.GatewaySuccessAddRequests.Add(1)
	return &emptypb.Empty{}, nil
}