package server

import (
	"homework/internal/storage/facade"
	pb "homework/pkg/api"
)

type implementation struct {
	pb.UnimplementedStorageServer

	storage facade.StorageFacade
}

func New(storage facade.StorageFacade) pb.StorageServer {
	return &implementation{
		storage: storage,
	}
}
