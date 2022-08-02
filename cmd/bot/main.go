package main

import (
	servicePkg "homework/internal/service"
	storagePkg "homework/internal/storage"
	"log"
)

func main() {
	log.Println("start main")
	storage := storagePkg.NewStorage()

	tgService, err := servicePkg.NewService(storage)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := tgService.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go runREST()
	runGRPC(storage)
}
