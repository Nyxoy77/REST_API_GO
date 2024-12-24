package caching

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitializeRedis(addr string, password string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	resp, err := RedisClient.Ping(context.Background()).Result()
	fmt.Println(resp)
	if err != nil {
		log.Fatalf("An error occured starting the redis server")
		return
	}
	log.Println("Redis Running")
}
