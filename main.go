package main

import (
	"log"
	"profile/api"
)

func main() {
	port := 3000
	address := "127.0.0.1"

	server := api.New(address, port)

	log.Fatal(server.Run())
}
