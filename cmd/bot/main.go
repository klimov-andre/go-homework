package main

import (
	servicePkg "homework/internal/service"
	storagePkg "homework/internal/storage"
	"log"
)

func main() {
	log.Println("start main")
	var storage = storagePkg.NewStorage()

	var tgService *servicePkg.Service
	{
		s, err := servicePkg.NewService(storage)
		if err != nil {
			log.Fatal(err)
		}
		tgService = s
	}

	go func() {
		if err := tgService.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	runGRPC(storage)
}
