package main

import (
	"log"

	"github.com/Nyxoy/restAPI/caching"
	"github.com/Nyxoy/restAPI/cmd/api"
	"github.com/Nyxoy/restAPI/config"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	server := api.NewAPIServer(viper.GetString("SERVER_PORT"))
	caching.InitializeRedis("localhost:6379", "", 0)
	if error := server.Run(); error != nil {
		log.Fatalf("An error occured at the time of  initialisation %s", error)
	}

}
