package server

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage/facade/mocks"
	"homework/internal/storage/models"
	pb "homework/pkg/api/gateway"
	"testing"
)

func TestGatewayServer_MovieUpdate(t *testing.T) {
	t.Run("incorrect input", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		tests := []struct {
			name string
			argMovie *pb.Movie
			expectedCode codes.Code
		} {
			{
				"id",
				&pb.Movie{
					Id: 0,
				},
				codes.InvalidArgument,
			},
			{
				"year",
				&pb.Movie{
					Id: 1,
					Year: 9,
					Title: "title",
				},
				codes.InvalidArgument,
			},
			{
				"title",
				&pb.Movie{
					Id: 1,
					Year: 2009,
					Title: "fT8NTyBiyxZvGOoTaVL8xQ9YJMcd2yyD6UHKl4bXumqLe6yU0Y2VirxNiE1pcDyt5Fo3TecTNvDjB8ZqKfJ4zFEXVfGpqYgPZr2GPOM0RQD1Pip17qllUqL3qvxM10jHRasYjK7VrCSmITTER5suK2USeXdXTSd1jP7moNfOm0owThLtE4yu1cVjwoZzukXiWNDcrZn5VeTppeB9vldRdDgLxA6iE5eva8UrEzvdJb9HGgLKlOtXrWYRw4nwUtBZ",
				},
				codes.InvalidArgument,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, err := server.MovieUpdate(context.Background(), &pb.GatewayMovieUpdateRequest{
					Movie: &pb.Movie{
						Id: 0,
					},
				})

				assert.Error(t, err)
				assert.Equal(t, status.Code(err), test.expectedCode)
				storage.AssertNotCalled(t, "Update")
			})
		}
	})

	t.Run("valid input success update", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("Update", mock.Anything, uint64(1), &models.Movie{
			Title: "title",
			Year:  2000,
		}).Return(nil)

		_, err := server.MovieUpdate(context.Background(), &pb.GatewayMovieUpdateRequest{
			Movie: &pb.Movie{
				Id:    1,
				Title: "title",
				Year:  2000,
			},
		})

		assert.NoError(t, err)
		storage.AssertNumberOfCalls(t, "Update", 1)
	})

	t.Run("valid input unsuccessful update", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("Update", mock.Anything, uint64(1), &models.Movie{
			Title: "title",
			Year:  2000,
		}).Return(errors.New("error"))

		_, err := server.MovieUpdate(context.Background(), &pb.GatewayMovieUpdateRequest{
			Movie: &pb.Movie{
				Id:    1,
				Title: "title",
				Year:  2000,
			},
		})

		assert.Error(t, err)
		storage.AssertNumberOfCalls(t, "Update", 1)
	})
}