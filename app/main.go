package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codecrafters-io/redis-starter-go/app/server"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "6379", "port on which to run the server")

	flag.Parse()

	server, err := server.NewServer(fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	server.Serve()
}
