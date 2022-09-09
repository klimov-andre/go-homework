package main

import (
	"github.com/sirupsen/logrus"
	"homework/config/kafka"
	"homework/internal/api/gateway/kafka/sender"
	storagePkg "homework/internal/api/storage/client"
	_ "homework/internal/logs"
	servicePkg "homework/internal/tgservice"
	"homework/pkg/tracing/exporter"
	"net/http"
	"os"
)

func main() {
	logrus.Info("start gateway service")

	exp, err := exporter.NewExporter(os.Stdout, "gateway", "0.0.6")
	if err != nil {
		logrus.Fatal(err)
	}
	exp.Run()
	defer exp.DownExporter()

	storage, err := storagePkg.New()
	if err != nil {
		logrus.Fatal(err)
	}

	tgService, err := servicePkg.NewService(storage)
	if err != nil {
		logrus.Fatal(err)
	}

	kafkaSender, err := sender.NewSender(kafka.Brokers)
	if err != nil {
		logrus.Fatal(err)
	}

	go func() {
		if err := tgService.Run(); err != nil {
			logrus.Fatal(err)
		}
	}()

	go http.ListenAndServe(":6000", http.DefaultServeMux)
	go runREST()
	runGRPC(storage, kafkaSender)
}
