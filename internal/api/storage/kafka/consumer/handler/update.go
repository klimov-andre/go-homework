package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
	"homework/config/storage"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (h Handler) Update(ctx context.Context, key, value []byte) {
	var span trace.Span
	ctx, span = otel.Tracer(storage.SpanTraceName).Start(ctx, "Update")
	defer span.End()

	log := logrus.WithField("request", "update")

	request := &pb.StorageMovieUpdateRequest{}
	if err := proto.Unmarshal(value, request); err != nil {
		span.RecordError(err)
		log.Errorf("value unmarshal error %v", err)
		return
	}

	log.Debugf("request %v, key=%v", request.String(), string(key))

	m := request.GetMovie()
	err := h.storage.Update(context.Background(), m.GetId(),&models.Movie{
		Title: m.GetTitle(),
		Year:  int(m.GetYear()),
	})
	if err != nil {
		span.RecordError(err)
		log.Warningf("movie update error %v", err)
		return
	}
}
