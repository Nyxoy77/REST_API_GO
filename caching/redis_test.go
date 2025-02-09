package caching

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRedisServer(t *testing.T) {
	err := InitializeRedis("localhost:6379", "", 0)
	assert.NoError(t, err, "Intialising redis should not return any error")
	resp, err2 := RedisClient.Ping(context.Background()).Result()
	assert.NoError(t, err2, "The ping should not return an error")
	assert.Equal(t, "PONG", resp, "Expected pong response from redis")
	_ = RedisClient.Close()
}
func TestAvailability(t *testing.T) {
	redisLocal := viper.GetString("REDIS_SERVER")
	client := redis.NewClient(&redis.Options{
		Addr: redisLocal,
	})
	_, err := client.Ping(context.Background()).Result()
	assert.NoError(t, err, "Redis server should be available ")
	_ = client.Close()
}
