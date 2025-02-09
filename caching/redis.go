package caching

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitializeRedis(addr string, password string, db int) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("An error occured starting the redis server")
		return fmt.Errorf("An error occured starting the redis server %w", err)
	}
	log.Println("Redis Running")
	return nil
}
