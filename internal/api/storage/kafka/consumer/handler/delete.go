package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	pb "homework/pkg/api/storage"
)

func (h Handler) Delete(key, value []byte) {
	log := logrus.WithField("request", "delete")
	request := &pb.StorageMovieDeleteRequest{}

	if err := proto.Unmarshal(value, request); err != nil {
		log.Errorf("value unmarshal error %v", err)
		return
	}

	log.Debugf("request %v, key=%v", request.String(), string(key))

	if err := h.storage.Delete(context.Background(), request.GetId()); err != nil {
		log.Warningf("delete movie error %v", err)
		return
	}
}
