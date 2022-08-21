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

func TestGatewayServer_MovieCreate(t *testing.T) {
	t.Run("incorrect input", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		tests := []struct {
			name string
			expectedCode codes.Code
			title string
			year int32
		} {
			{
				"year",
				codes.InvalidArgument,
				"Hello",
				1000,
			},
			{
				"title",
				codes.InvalidArgument,
				"MxgII9AQolAP8aYX9eqxkgXIybQkPziQ3g8AnYol0KjGXafENjpfNncDVZkaRfRCARJJxzvGBiSxxtWU27YQh0uOnvIlXUqGgyoE8EJ5R6K6bNvUBTYsoemoTX7KrgRKlhNcBfhiL8XSrtmw8AUKl1z4VSZS7ZK9BKatSF5W9mptXpqqu01fIs1gNeO1qxc88f7iy2eD7kk5LzsC0nNBWMvnEDZk9R6aBEENOK7Q27BKfK24vdDQwbiabiECNh9E",
				1999,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, err := server.MovieCreate(context.Background(), &pb.GatewayMovieCreateRequest{
					Title: test.title,
					Year: test.year,
				})

				assert.Equal(t, status.Code(err), test.expectedCode)
				storage.AssertNotCalled(t, "Add")
			})
		}
	})

	t.Run("valid input success add", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("Add", mock.Anything, &models.Movie{
			Title: "Hello",
			Year:  1999,
		}).Return(uint64(1), nil)

		_, err := server.MovieCreate(context.Background(), &pb.GatewayMovieCreateRequest{
			Title: "Hello",
			Year: 1999,
		})

		assert.NoError(t, err)
		storage.AssertNumberOfCalls(t, "Add", 1)
	})

	t.Run("valid input unsuccessful add", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("Add", mock.Anything, &models.Movie{
			Title: "Hello",
			Year:  1999,
		}).Return(uint64(0), errors.New("error"))

		_, err := server.MovieCreate(context.Background(), &pb.GatewayMovieCreateRequest{
			Title: "Hello",
			Year: 1999,
		})

		assert.Error(t, err)
		storage.AssertNumberOfCalls(t, "Add", 1)
	})
}
