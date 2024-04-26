package main

import (
	"log"

	"github.com/codecrafters-io/redis-starter-go/app/server"
)

func main() {
	server, err := server.NewServer("0.0.0.0:6379")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	server.Serve()
}
