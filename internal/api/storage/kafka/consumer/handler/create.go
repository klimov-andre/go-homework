package handler

import (
	"context"
	"google.golang.org/protobuf/proto"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
	"log"
)

func (h Handler) Create(key, value []byte) {
	request := &pb.StorageMovieCreateRequest{}

	err := proto.Unmarshal(value, request)
	if err != nil {
		// TODO
	}
	log.Printf("create request %v %v", request.String(), string(key))

	id, err := h.storage.Add(context.Background(), &models.Movie{
		Title: request.Title,
		Year:  int(request.Year),
	})

	_ = id // TODO
}
