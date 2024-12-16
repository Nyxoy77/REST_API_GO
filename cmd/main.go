package main

import (
	"log"

	"github.com/Nyxoy/restAPI/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	if error := server.Run(); error != nil {
		log.Fatal("An error occured at the time of  initialisation")
	}

}
