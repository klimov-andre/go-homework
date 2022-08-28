package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieCreateQueued(context.Context, *pb.GatewayMovieCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieCreateQueued not implemented")
}