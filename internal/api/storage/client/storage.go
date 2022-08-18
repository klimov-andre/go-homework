package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	storageCfg "homework/config/storage"
	"homework/internal/storage/facade"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
	"io"
)

var _ facade.StorageFacade = (*grpcStorage)(nil)

// grpcStorage implements StorageFacade and use remote Storage service
type grpcStorage struct {
	storage pb.StorageClient
}

func New() (facade.StorageFacade, error) {
	connection, err := grpc.Dial(storageCfg.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcStorage{
		storage: pb.NewStorageClient(connection),
	}, nil
}

func (s *grpcStorage) List(ctx context.Context, limit, offset int, order string) ([]*models.Movie, error) {
	request := &pb.StorageMovieListRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
		Order: order,
	}
	stream, err := s.storage.MovieList(ctx, request)
	if err != nil {
		return nil, err
	}

	var movies []*models.Movie
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			return movies, nil
		}

		for _, m := range response.Movie {
			movies = append(movies, models.MovieFromPb(m))
		}
	}
}

func (s *grpcStorage) Add(ctx context.Context, m *models.Movie) error {
	request := &pb.StorageMovieCreateRequest{
		Title: m.Title,
		Year: int32(m.Year),
	}

	_, err := s.storage.MovieCreate(ctx, request)
	return err
}

func (s *grpcStorage) Update(ctx context.Context, id uint64, newMovie *models.Movie) error {
	request := &pb.StorageMovieUpdateRequest{
		Movie: &pb.Movie{
			Id:    id,
			Title: newMovie.Title,
			Year: int32(newMovie.Year),
		},
	}

	_, err := s.storage.MovieUpdate(ctx, request)
	return err
}

func (s *grpcStorage) Delete(ctx context.Context, id uint64) error {
	request := &pb.StorageMovieDeleteRequest{
		Id: id,
	}
	_, err := s.storage.MovieDelete(ctx, request)
	return err
}

func (s *grpcStorage) GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error) {
	request := &pb.StorageMovieGetOneRequest{
		Id: id,
	}
	response, err := s.storage.MovieGetOne(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.MovieFromPb(response.Movie), nil
}