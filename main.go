package main

import "github.com/nilspolek/grpcCalculator/server"

func main() {
	server := server.Server{}
	if err := server.Start(); err != nil {
		panic(err)
	}
}

