package server

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"homework/internal/storage/facade/mocks"
	pb "homework/pkg/api/gateway"
	"testing"
)

func TestGatewayServer_MovieDelete(t *testing.T) {
	t.Run("incorrect input", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		_, err := server.MovieDelete(context.Background(), &pb.GatewayMovieDeleteRequest{
			Id: 0,
		})

		assert.Error(t, err)
		assert.Equal(t, status.Code(err), codes.InvalidArgument)
		storage.AssertNotCalled(t, "Delete")
	})

	t.Run("valid input success delete", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("Delete", mock.Anything, uint64(1)).
			Return(nil)

		_, err := server.MovieDelete(context.Background(), &pb.GatewayMovieDeleteRequest{
			Id: 1,
		})

		assert.NoError(t, err)
		storage.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("valid input unsuccessful end", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("Delete", mock.Anything, uint64(1)).
			Return(errors.New("error"))

		_, err := server.MovieDelete(context.Background(), &pb.GatewayMovieDeleteRequest{
			Id: 1,
		})

		assert.Error(t, err)
		storage.AssertNumberOfCalls(t, "Delete", 1)
	})
}