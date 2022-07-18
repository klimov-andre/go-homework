package main

import (
	"homework/internal/robot"
	"log"
)

func main() {
	log.Println("start main")
	cmd, err := robot.Init()
	if err != nil {
		log.Panic(err)
	}

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
