package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework/config"
	pb "homework/pkg/api"
	"log"
	"time"
)

func main() {
	conns, err := grpc.Dial(config.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewAdminClient(conns)

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 30)
		list, errG := client.MovieList(ctx, &pb.MovieListRequest{})
		if errG != nil {
			log.Fatal(err)
		}

		fmt.Println(list)
	}
}
