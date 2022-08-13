package server

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/internal/storage/connections"
	pb "homework/pkg/api"
)

func (i *gatewayServer) MovieDelete(ctx context.Context, req *pb.MovieDeleteRequest) (*emptypb.Empty, error) {
	if err := i.storage.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if errors.Is(err, connections.ErrTimeout) {
		return nil, status.Error(codes.DeadlineExceeded, err.Error())
	}

	return &emptypb.Empty{}, nil
}
