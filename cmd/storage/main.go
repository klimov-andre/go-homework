package main

import (
	"google.golang.org/grpc"
	"homework/config/kafka"
	storageCfg "homework/config/storage"
	"homework/internal/api/storage/kafka/consumer"
	"homework/internal/api/storage/kafka/consumer/handler"
	apiPkg "homework/internal/api/storage/server"
	"homework/internal/storage/facade"
	pb "homework/pkg/api/storage"
	"log"
	"net"
)

func runGRPC(storage facade.StorageFacade) {
	listener, err := net.Listen(storageCfg.GrpcProtocol, storageCfg.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStorageServer(grpcServer, apiPkg.New(storage))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func runKafkaHandler(storage facade.StorageFacade) {
	cons, err := consumer.NewConsumer(kafka.Brokers, handler.NewHandler(storage))
	if err != nil {
		log.Fatal(err)
	}

	if err = cons.Serve(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("start storage service")
	storage, err := facade.NewStorage(storageCfg.DbDSN)
	if err != nil {
		log.Fatal(err)
	}

	runKafkaHandler(storage)
	runGRPC(storage)
}
