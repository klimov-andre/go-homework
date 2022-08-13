package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "homework/pkg/api"
)

func (i *gatewayServer) MovieDelete(ctx context.Context, req *pb.MovieDeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	if id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id must be > 0")
	}

	if err := i.storage.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
