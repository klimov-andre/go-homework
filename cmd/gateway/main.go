package main

import (
	"homework/config/gateway"
	"homework/internal/api/gateway/kafka/sender"
	storagePkg "homework/internal/api/storage/client"
	servicePkg "homework/internal/tgservice"
	"log"
)

func main() {
	log.Println("start gateway service")

	storage, err := storagePkg.New()
	if err != nil {
		log.Fatal(err)
	}

	tgService, err := servicePkg.NewService(storage)
	if err != nil {
		log.Fatal(err)
	}

	kafkaSender, err := sender.NewSender(gateway.Brokers)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := tgService.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go runREST()
	runGRPC(storage, kafkaSender)
}
