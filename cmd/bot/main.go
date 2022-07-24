package main

import (
	"homework/internal/service"
	"log"
)

func main() {
	log.Println("start main")
	service, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
