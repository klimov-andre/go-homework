package main

import (
	"github.com/sirupsen/logrus"
	"homework/config/kafka"
	"homework/internal/api/gateway/kafka/sender"
	storagePkg "homework/internal/api/storage/client"
	_ "homework/internal/logs"
	servicePkg "homework/internal/tgservice"
	"net/http"
)

func main() {
	logrus.Info("start gateway service")

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
