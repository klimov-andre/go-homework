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

func (g *gatewayServer) MovieDeleteQueued(ctx context.Context, req *pb.GatewayMovieDeleteRequest) (*emptypb.Empty, error) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "MovieDeleteQueued")
	defer span.End()

	metrics.GatewayTotalDeleteRequests.Add(1)

	id := req.GetId()
	request := &pbStorage.StorageMovieDeleteRequest{
		Id: id,
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		span.RecordError(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	g.kafkaSender.SendMessage(kafka.TopicDelete, fmt.Sprint(id), msg)
	return &emptypb.Empty{}, nil
}