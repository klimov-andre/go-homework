package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/pkg/api"
)

func (i *gatewayServer) MovieList(ctx context.Context, req *pb.GatewayMovieListRequest) (*pb.MovieListResponse, error) {
	limit := int(req.GetLimit())
	if limit <= 0 {
		return nil, status.Error(codes.InvalidArgument, "limit must be > 0")
	}

	order := "ASC"
	switch req.GetOrder() {
	case pb.ListOrder_LIST_ORDER_DESC:
		order = "DESC"
	}

	list, err := i.storage.List(ctx, limit, 0, order)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Movie, 0, len(list))
	for _, m := range list {
		result = append(result, &pb.Movie{
			Id:    m.Id,
			Title: m.Title,
			Year:  int32(m.Year),
		})
	}

	return &pb.MovieListResponse{
		Movie: result,
	}, nil
}