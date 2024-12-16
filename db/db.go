package db

import (
	"github.com/go-resty/resty/v2"
)

func CreateRestyClient() *resty.Client {
	// Initialize Resty Client
	client := resty.New()

	apiKey := API // Replace with your actual API key

	client.SetHeader("apikey", apiKey)
	client.SetHeader("Authorization", "Bearer "+apiKey)
	client.SetHeader("Accept", "application/json")
	return client
}
