package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"homework/internal/storage/models"
	pb "homework/pkg/api/storage"
	"io"
	"sort"
	"testing"
)

func Test_Integration_MovieCreate(t *testing.T){
	listener := runTestStorage(t)

	storage := setupClientStorage(t, listener)
	defer storage.TearDown()

	tests := []struct {
		name string
		request *pb.StorageMovieCreateRequest
		wantError bool
	} {
		{
			"success",
			&pb.StorageMovieCreateRequest{
				Title: "Hello",
				Year: 2000,
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := storage.client.MovieCreate(context.Background(), test.request)

			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, response.GetId()>0)
			}
		})
	}
}

func Test_Integration_MovieDelete(t *testing.T){
	listener := runTestStorage(t)
	storage := setupClientStorage(t, listener)
	defer storage.TearDown()

	id, err := oneMoviePrepare(&models.Movie{
		Title: "testmovie",
		Year:  2000,
	}, storage)
	if err != nil {
		t.Fatalf("could not prepare movie: %v", err.Error())
	}
	tests := []struct {
		name string
		request *pb.StorageMovieDeleteRequest
		wantError bool
	} {
		{
			"success delete",
			&pb.StorageMovieDeleteRequest{
				Id: id,
			},
			false,
		},
		{
			"delete nonexistent",
			&pb.StorageMovieDeleteRequest{
				Id: 10000,
			},
			false,

		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err = storage.client.MovieDelete(context.Background(), test.request)

			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Integration_MovieUpdate(t *testing.T){
	listener := runTestStorage(t)
	storage := setupClientStorage(t, listener)
	defer storage.TearDown()

	id, err := oneMoviePrepare(&models.Movie{
		Title: "testmovie",
		Year:  2000,
	}, storage)
	if err != nil {
		t.Fatalf("could not prepare movie: %v", err.Error())
	}
	tests := []struct {
		name string
		request *pb.StorageMovieUpdateRequest
		wantError bool
	} {
		{
			"success update",
			&pb.StorageMovieUpdateRequest{
				Movie: &pb.Movie{
					Id:    id,
					Title: "new test title",
					Year:  2000,
				},
			},
			false,
		},
		{
			"update nonexistent",
			&pb.StorageMovieUpdateRequest{
				Movie: &pb.Movie{
					Id:    100000,
					Title: "new test title",
					Year:  2000,
				},
			},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err = storage.client.MovieUpdate(context.Background(), test.request)

			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Integration_MovieGetOne(t *testing.T){
	listener := runTestStorage(t)
	storage := setupClientStorage(t, listener)
	defer storage.TearDown()

	id, err := oneMoviePrepare(&models.Movie{
		Title: "testmovie",
		Year:  2000,
	}, storage)
	if err != nil {
		t.Fatalf("could not prepare movie: %v", err.Error())
	}
	tests := []struct {
		name string
		request *pb.StorageMovieGetOneRequest
		wantError bool
		wantMovie *pb.Movie
	} {
		{
			"success update",
			&pb.StorageMovieGetOneRequest{
				Id: id,
			},
			false,
			&pb.Movie{
				Id:    id,
				Title: "testmovie",
				Year:  2000,
			},
		},
		{
			"get nonexistent",
			&pb.StorageMovieGetOneRequest{
				Id: 10000,
			},
			true,
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := storage.client.MovieGetOne(context.Background(), test.request)

			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.Equal(t,  test.wantMovie, response.GetMovie())
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Integration_MovieList(t *testing.T){
	listener := runTestStorage(t)
	storage := setupClientStorage(t, listener)
	defer storage.TearDown()

	movies := []models.Movie{
		{
			Title: "movie 1",
			Year: 2001,
		},
		{
			Title: "movie 2",
			Year: 2002,
		},
		{
			Title: "movie 3",
			Year: 2003,
		},
	}
	_, err := movieListPrepare(storage, movies...)
	if err != nil {
		t.Fatalf("could not prepare movies: %v", err.Error())
	}

	tests := []struct {
		name string
		request *pb.StorageMovieListRequest
	} {
		{
			"success list by 1 movie",
			&pb.StorageMovieListRequest{
				Limit: 1,
				Offset: 0,
				Order: "ASC",
			},
		},
		{
			"success list by 2 movie",
			&pb.StorageMovieListRequest{
				Limit: 2,
				Offset: 0,
				Order: "ASC",
			},
		},
		{
			"list all movies",
			&pb.StorageMovieListRequest{
				Limit: 100,
				Offset: 0,
				Order: "ASC",
			},
		},
		{
			"list movies descended",
			&pb.StorageMovieListRequest{
				Limit: 100,
				Offset: 0,
				Order: "DESC",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := storage.client.MovieList(context.Background(), test.request)
			assert.NoError(t, err)
			var allMovies []*models.Movie
			for {
				r, err := response.Recv()
				if err == io.EOF {
					sorted := false
					// check if slice is sorted correctly
					if test.request.GetOrder() == "ASC" {
						sorted = sort.SliceIsSorted(allMovies, func(p, q int) bool { return allMovies[p].Id < allMovies[q].Id })
					} else {
						sorted = sort.SliceIsSorted(allMovies, func(p, q int) bool { return allMovies[p].Id > allMovies[q].Id })
					}
					assert.True(t, sorted)
					break
				}

				// check limits for one stream request
				assert.True(t, len(r.Movie) <= int(test.request.Limit))

				for _, m := range r.GetMovie() {
					allMovies = append(allMovies, models.MovieFromPb(m))
				}
			}
		})
	}
}