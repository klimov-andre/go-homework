package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework/config"
	pb "homework/pkg/api"
	"log"
	"sync"
)

func main() {
	conns, err := grpc.Dial(config.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewAdminClient(conns)

	var wg = sync.WaitGroup{}
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			list, errG := client.MovieList(ctx, &pb.MovieListRequest{})
			if errG != nil {
				fmt.Printf("%d %v\n", k, errG.Error())
				return
			}

			fmt.Printf("%d %v\n", k, list)
		}(i)

	}

	wg.Wait()
}
