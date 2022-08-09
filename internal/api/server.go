package api

import (
	"homework/internal/storage/facade"
	pb "homework/pkg/api"
)

type implementation struct {
	pb.UnimplementedAdminServer

	storage facade.StorageFacade
}

func New(storage facade.StorageFacade) pb.AdminServer {
	return &implementation{
		storage: storage,
	}
}
