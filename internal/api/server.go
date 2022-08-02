package api

import (
	storagePkg "homework/internal/storage"
	pb "homework/pkg/api"
)

type implementation struct {
	pb.UnimplementedAdminServer

	storage *storagePkg.Storage
}

func New(storage *storagePkg.Storage) pb.AdminServer {
	return &implementation{
		storage: storage,
	}
}
