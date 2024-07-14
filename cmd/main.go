package main

import (
	"log"
	"github.com/JuanMartinez503/Golang-Backend/cmd/api"
)
func main() {
	// Run the server
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	
	
}