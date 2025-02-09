package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("cmd")
	// viper.AddConfigPath(".")
	if os.Getenv("TEST_MODE") == "true" {
		viper.SetConfigFile("C:/Users/Shivam Rai/OneDrive/Desktop/DESKTOP/go_lco/Project/RESTAPI/cmd/.env")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("There was an error loading the env file : ", err)
		return fmt.Errorf("there was an error loading the env file : %w", err)
	}
	viper.SetDefault("SERVER_PORT", ":8081")
	return nil
}
