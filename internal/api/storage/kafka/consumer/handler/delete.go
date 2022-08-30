package handler

import (
	"context"
	"google.golang.org/protobuf/proto"
	pb "homework/pkg/api/storage"
	"log"
)

func (h Handler) Delete(key, value []byte) {
	request := &pb.StorageMovieDeleteRequest{}

	err := proto.Unmarshal(value, request)
	if err != nil {
		// TODO
	}
	log.Printf("delete request %v %v", request.String(), string(key))

	err = h.storage.Delete(context.Background(), request.GetId())
	_ = err // TODO
}
