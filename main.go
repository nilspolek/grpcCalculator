package main

import (
	"log"

	"github.com/nilspolek/grpcCalculator/server"
)

func main() {
	server := server.Server{Address: ":50051"}
	log.Println("Starting server on", server.Address)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
