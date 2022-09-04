package server

import (
	"homework/internal/api/storage/publisher"
	"homework/internal/storage/facade"
	pb "homework/pkg/api/storage"
)

type storageServer struct {
	pb.UnimplementedStorageServer

	storage facade.StorageFacade
	publisher *publisher.Publisher
}

func New(storage facade.StorageFacade) pb.StorageServer {
	return &storageServer{
		storage: storage,
		publisher: publisher.NewRedisPublisher(),
	}
}
