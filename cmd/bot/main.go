package main

import (
	servicePkg "homework/internal/service"
	"homework/internal/storage/facade"
	"log"
)

func main() {
	log.Println("start main")

	storage, err := facade.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

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
