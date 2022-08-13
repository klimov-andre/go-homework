package server

import (
	"homework/internal/storage/facade"
	pb "homework/pkg/api"
)

type storageServer struct {
	pb.UnimplementedStorageServer

	storage facade.StorageFacade
}

func New(storage facade.StorageFacade) pb.StorageServer {
	return &storageServer{
		storage: storage,
	}
}
