package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework/internal/api/gateway/subscriber"
	"homework/pkg/api/gateway"
)

func (g *gatewayServer) MovieSubscribe(_ *emptypb.Empty, srv gateway.Gateway_MovieSubscribeServer) error {
	// create new redis subscriber for current request
	subscriber := subscriber.NewRedisSubscriber()
	go subscriber.ReceiveMessages()

loop:
	for  {
		select {
		case <-srv.Context().Done():
			subscriber.Done()
			logrus.Debugf("done subscribe request")
			break loop
		case m := <- subscriber.Channel():
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
