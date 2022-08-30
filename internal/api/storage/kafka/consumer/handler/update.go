package handler

import (
	"context"
	"google.golang.org/protobuf/proto"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
	"log"
)

func (h Handler) Update(key, value []byte) {
	request := &pb.StorageMovieUpdateRequest{}

	err := proto.Unmarshal(value, request)
	if err != nil {
		// TODO
		return
	}

	log.Printf("update request %v %v", request.String(), string(key))

	m := request.GetMovie()

	err = h.storage.Update(context.Background(), m.GetId(),&models.Movie{
		Title: m.GetTitle(),
		Year:  int(m.GetYear()),
	})

	_ = err // TODO
}
