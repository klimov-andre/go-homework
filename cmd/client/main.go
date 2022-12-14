package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework/config/gateway"
	pb "homework/pkg/api/gateway"
)

func main() {
	conns, err := grpc.Dial(gateway.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Panic(err)
	}

	client := pb.NewGatewayClient(conns)

	//var wg = sync.WaitGroup{}
	ctx := context.Background()
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go func(k int) {
	//		defer wg.Done()
			list, errG := client.MovieList(ctx, &pb.GatewayMovieListRequest{Limit: 1, Order: 2})
			if errG != nil {
				fmt.Printf("%d %v\n", 1, errG.Error())
				return
			}

		fmt.Printf("%d %v\n",1, list)
		//}(i)

	//}

	//wg.Wait()
}
