package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
	"homework/config/storage"
	pb "homework/pkg/api/storage"
)

func (h Handler) Delete(ctx context.Context, key, value []byte) {
	var span trace.Span
	ctx, span = otel.Tracer(storage.SpanTraceName).Start( ctx, "Delete")
	defer span.End()

	log := logrus.WithField("request", "delete")
	request := &pb.StorageMovieDeleteRequest{}

	if err := proto.Unmarshal(value, request); err != nil {
		span.RecordError(err)
		log.Errorf("value unmarshal error %v", err)
		return
	}

	log.Debugf("request %v, key=%v", request.String(), string(key))

	if err := h.storage.Delete(context.Background(), request.GetId()); err != nil {
		span.RecordError(err)
		log.Warningf("delete movie error %v", err)
		return
	}
}
