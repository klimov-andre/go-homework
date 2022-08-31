package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"homework/config/kafka"
	storageCfg "homework/config/storage"
	"homework/internal/api/storage/kafka/consumer"
	"homework/internal/api/storage/kafka/consumer/handler"
	apiPkg "homework/internal/api/storage/server"
	"homework/internal/storage/facade"
	pb "homework/pkg/api/storage"
	"homework/pkg/tracing/exporter"
	"net"
	"os"
)

func runGRPC(storage facade.StorageFacade) {
	listener, err := net.Listen(storageCfg.GrpcProtocol, storageCfg.GrpcPort)
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStorageServer(grpcServer, apiPkg.New(storage))

	if err = grpcServer.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}

func runKafkaHandler(storage facade.StorageFacade) {
	cons, err := consumer.NewConsumer(kafka.Brokers, handler.NewHandler(storage))
	if err != nil {
		logrus.Fatal(err)
	}

	if err = cons.Serve(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	logrus.Println("start storage service")

	exp, err := exporter.NewExporter(os.Stdout, "gateway", "0.0.6")
	if err != nil {
		logrus.Fatal(err)
	}
	exp.Run()
	defer exp.DownExporter()

	storage, err := facade.NewStorage(storageCfg.DbDSN)
	if err != nil {
		logrus.Fatal(err)
	}

	runKafkaHandler(storage)
	runGRPC(storage)
}
