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

func (h Handler) Create(ctx context.Context, key, value []byte) {
	var span trace.Span
	ctx, span = otel.Tracer(storage.SpanTraceName).Start( ctx, "Create")
	defer span.End()

	log := logrus.WithField("request", "create")

	request := &pb.StorageMovieCreateRequest{}
	if err := proto.Unmarshal(value, request); err != nil {
		span.RecordError(err)
		log.Errorf("value unmarshal error %v", err)
		return
	}

	log.Debugf("request %v, key=%v", request.String(), string(key))

	id, err := h.storage.Add(ctx, &models.Movie{
		Title: request.Title,
		Year:  int(request.Year),
	})
	if err != nil {
		span.RecordError(err)
		log.Warningf("add movie error %v", err)
		return
	}
	log.Debugf("successfully add movie %v", id)
}
