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

func TestGatewayServer_MovieGetOne(t *testing.T) {
	t.Run("incorrect input", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		_, err := server.MovieGetOne(context.Background(), &pb.GatewayMovieGetOneRequest{
			Id: 0,
		})

		assert.Error(t, err)
		assert.Equal(t, status.Code(err), codes.InvalidArgument)
		storage.AssertNotCalled(t, "GetOneMovie")
	})

	t.Run("valid input success GetOneMovie", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}
		expectedMovie := &pb.Movie{
			Id:    1,
			Title: "title",
			Year: 2000,
		}

		storage.On("GetOneMovie", mock.Anything, uint64(1)).
			Return(&models.Movie{
			Id:    1,
			Title: "title",
			Year:  2000,
		}, nil)

		actualRequest, err := server.MovieGetOne(context.Background(), &pb.GatewayMovieGetOneRequest{
			Id: 1,
		})

		assert.Equal(t, actualRequest.Movie, expectedMovie)
		assert.NoError(t, err)
		storage.AssertNumberOfCalls(t, "GetOneMovie", 1)
	})

	t.Run("valid input unsuccessful end", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("GetOneMovie", mock.Anything, uint64(1)).
			Return(nil, errors.New("error"))

		_, err := server.MovieGetOne(context.Background(), &pb.GatewayMovieGetOneRequest{
			Id: 1,
		})

		assert.Error(t, err)
		storage.AssertNumberOfCalls(t, "GetOneMovie", 1)
	})
}