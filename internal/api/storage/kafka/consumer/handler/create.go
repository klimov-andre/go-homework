package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (h Handler) Create(key, value []byte) {
	log := logrus.WithField("request", "create")

	request := &pb.StorageMovieCreateRequest{}
	if err := proto.Unmarshal(value, request); err != nil {
		log.Errorf("value unmarshal error %v", err)
		return
	}

	log.Debugf("request %v, key=%v", request.String(), string(key))

	id, err := h.storage.Add(context.Background(), &models.Movie{
		Title: request.Title,
		Year:  int(request.Year),
	})
	if err != nil {
		log.Warningf("add movie error %v", err)
		return
	}
	log.Debugf("successfully add movie %v", id)
}
