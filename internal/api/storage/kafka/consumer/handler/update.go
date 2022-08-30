package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
)

func (h Handler) Update(key, value []byte) {
	log := logrus.WithField("request", "update")

	request := &pb.StorageMovieUpdateRequest{}
	if err := proto.Unmarshal(value, request); err != nil {
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
		log.Warningf("movie update error %v", err)
		return
	}
}
