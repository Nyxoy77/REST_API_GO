package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./cmd")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("There was an error loading the env file : ", err)
	}
	viper.SetDefault("SERVER_PORT", ":8081")
}
