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

func Test_orderToString(t *testing.T) {
	tests := []struct{
		name string
		arg pb.ListOrder
		want string
	} {
		{
			"default",
			pb.ListOrder_LIST_ORDER_UNSPECIFIED,
			"ASC",
		},
		{
			"asc",
			pb.ListOrder_LIST_ORDER_ASC,
			"ASC",
		},
		{
			"desc",
			pb.ListOrder_LIST_ORDER_DESC,
			"DESC",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := orderToString(test.arg)

			assert.Equal(t, actual, test.want)
		})
	}
}

func TestGatewayServer_MovieList(t *testing.T) {
	t.Run("incorrect input", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		_, err := server.MovieList(context.Background(), &pb.GatewayMovieListRequest{
			Limit: 0,
			Order: pb.ListOrder_LIST_ORDER_ASC,
		})

		assert.Error(t, err)
		assert.Equal(t, status.Code(err), codes.InvalidArgument)
		storage.AssertNotCalled(t, "List")
	})

	t.Run("valid input success list", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}
		expectedMovies := []*pb.Movie{
			{
				Id:    1,
				Title: "title",
				Year: 2000,
			},
		}

		storage.On("List", mock.Anything, 1, 0, "ASC").
			Return([]*models.Movie{
				{
					Id:    1,
					Title: "title",
					Year:  2000,
				},
			}, nil)

		actualRequest, err := server.MovieList(context.Background(), &pb.GatewayMovieListRequest{
			Limit: 1,
			Order: pb.ListOrder_LIST_ORDER_ASC,
		})

		assert.Equal(t, actualRequest.Movie, expectedMovies)
		assert.NoError(t, err)
		storage.AssertNumberOfCalls(t, "List", 1)
	})

	t.Run("valid input unsuccessful list", func(t *testing.T) {
		storage := &mocks.StorageFacade{}
		server := &gatewayServer{
			storage:  storage,
		}

		storage.On("List", mock.Anything, 2, 0, "ASC").
			Return(nil, errors.New("error"))

		_, err := server.MovieList(context.Background(), &pb.GatewayMovieListRequest{
			Limit: 2,
			Order: pb.ListOrder_LIST_ORDER_ASC,
		})

		assert.Error(t, err)
		storage.AssertNumberOfCalls(t, "List", 1)
	})
}