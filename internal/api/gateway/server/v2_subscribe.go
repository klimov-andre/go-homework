package server

import (
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieSubscribe(_ *emptypb.Empty, srv gateway.Gateway_MovieSubscribeServer) error {
	go g.subscriber.ReceiveMessages()

loop:
	for  {
		select {
		case <-srv.Context().Done():
			g.subscriber.Done()
			break loop
		case m := <- g.subscriber.Channel():
			srv.Send(&gateway.GatewayMovieSubscribeRequest{
				Movie: &gateway.Movie{
					Id:    m.Id,
					Title: m.Title,
					Year:  int32(m.Year),
				}})
		}
	}

	return nil
}
