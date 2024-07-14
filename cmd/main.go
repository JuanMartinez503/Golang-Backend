package main

import (
	"github.com/juanmartinez503/Golang-Backend/cmd/api"
	"log"
)
func main() {
	// Run the server
	server := api.NewAPIServer(":8080", nil)
	if err := server.run(); err != nil {
		log.Fatal(err)
	}
	
}