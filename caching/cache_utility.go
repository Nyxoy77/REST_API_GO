package caching

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var ctx = context.Background()

func SetCache(key string, value interface{}, ttl time.Duration) error {
	if RedisClient == nil {
		return fmt.Errorf("redis client is not initialized")
	}
	data, err := json.Marshal(value)
	if err != nil {
		log.Println("An error occured while setting the cache")
		return err
	}
	return RedisClient.Set(ctx, key, data, ttl).Err()
}

func BlackListToken(key string, ttl time.Duration) error {
	if RedisClient == nil {
		return fmt.Errorf("redis client is not initialized")
	}
	actKey := fmt.Sprintf("blacklist:%s", key)

	return RedisClient.Set(ctx, actKey, true, ttl).Err()
}

func GetCache(key string, dest interface{}) error {
	if RedisClient == nil {
		return fmt.Errorf("redis client is not initialized")
	}
	data, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		log.Println("An error occured while setting the cache")
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func DeleteCache(key string) error {
	if RedisClient == nil {
		return fmt.Errorf("redis client is not initialized")
	}
	_, err := RedisClient.Del(ctx, key).Result()
	if err != nil {
		log.Println("Error while deleteing the cache ")
		return err
	}
	log.Println("Deleted the cache successfully")
	return nil
}

func ExistCache(key string) bool {
	if RedisClient == nil {
		log.Println("Redis Client is not initialized ")
		return false
	}
	exists, _ := RedisClient.Exists(ctx, key).Result()

	if exists > 0 {
		return true
	}
	return false
}
