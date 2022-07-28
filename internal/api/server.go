package api

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (i *implementation) MovieCreate(_ context.Context, req *pb.MovieCreateRequest) (*pb.MovieCreateResponse, error) {
	if _, err := i.storage.Add(req.GetTitle(), int(req.GetYear())); err != nil {
		if errors.Is(err, storagePkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, storagePkg.ErrMovieExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MovieCreateResponse{}, nil
}

func (i *implementation) MovieList(_ context.Context, _ *pb.MovieListRequest) (*pb.MovieListResponse, error) {
	list := i.storage.List()
	result := make([]*pb.Movie, 0, len(list))
	for _, m := range list {
		result = append(result, &pb.Movie{
			Id:    m.Id(),
			Title: m.Title(),
			Year:  int32(m.Year()),
		})
	}

	return &pb.MovieListResponse{
		Movie: result,
	}, nil
}

func (i *implementation) MovieUpdate(_ context.Context, req *pb.MovieUpdateRequest) (*pb.MovieUpdateResponse, error) {
	m := req.GetMovie()
	if _, err := i.storage.Update(m.GetId(), m.GetTitle(), int(m.GetYear())); err != nil {
		if errors.Is(err, storagePkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, storagePkg.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MovieUpdateResponse{}, nil
}

func (i *implementation) MovieDelete(_ context.Context, req *pb.MovieDeleteRequest) (*pb.MovieDeleteResponse, error) {
	if err := i.storage.Delete(req.GetId()); err != nil {
		if errors.Is(err, storagePkg.ErrMovieNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.MovieDeleteResponse{}, nil
}
