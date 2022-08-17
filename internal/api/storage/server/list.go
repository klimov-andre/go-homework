package server

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage/connections"
	api "homework/pkg/api/include"
	pb "homework/pkg/api/storage"
)

func (i *storageServer) MovieList(req *pb.StorageMovieListRequest, srv pb.Storage_MovieListServer) error {
	limit, order := int(req.GetLimit()), req.GetOrder()
	for offset := int(req.GetOffset());; offset += limit {
		list, err := i.storage.List(srv.Context(), limit, offset, order)
		if err != nil {
			if errors.Is(err, connections.ErrTimeout) {
				return status.Error(codes.DeadlineExceeded, err.Error())
			}
			return status.Error(codes.Internal, err.Error())
		}

		// exit when no other movies in DB
		if len(list) == 0 {
			return nil
		}

		result := make([]*api.Movie, 0, len(list))
		for _, m := range list {
			result = append(result, &api.Movie{
				Id:    m.Id,
				Title: m.Title,
				Year:  int32(m.Year),
			})
		}

		srv.Send(&pb.StorageMovieListResponse{
			Movie: result,
		})
	}
}