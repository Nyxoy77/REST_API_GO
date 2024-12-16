package db

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func CreateRestyClient() *resty.Client {
	// Initialize Resty Client
	client := resty.New()

	apiKey := viper.GetString("DB_PASS") // Replace with your actual API key

	client.SetHeader("apikey", apiKey)
	client.SetHeader("Authorization", "Bearer "+apiKey)
	client.SetHeader("Accept", "application/json")
	return client
}
