package main

import (
	"log"
	"projectgo/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("error occured while starting api server: %v", err)
	}
}
