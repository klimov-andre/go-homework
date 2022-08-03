package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"homework/config"
	servicePkg "homework/internal/service"
	storagePkg "homework/internal/storage"
	"log"
)

func main() {
	log.Println("start main")
	conn, err := pgx.Connect(context.Background(), config.DbDSN)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer conn.Close(context.Background())

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
