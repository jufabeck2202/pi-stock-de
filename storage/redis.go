package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// RedisConnection func for connect to Redis server.
func GetRedisConnection() (*redis.Client, error) {
	// Define Redis database number.
	if redisClient != nil {
		return redisClient, nil
	}
	// Set Redis options.
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	// Connect to Redis server.
	fmt.Println("Connecting to Redis server...")
	newClient := redis.NewClient(options)

	redisClient = newClient
	return redisClient, nil
}

func Set(key string, value interface{}, ttl time.Duration) error {
	c, err := GetRedisConnection()
	if err != nil {
		return err
	}
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(context.Background(), key, p, ttl).Err()
}

func Get(key string, dest interface{}) error {
	c, err := GetRedisConnection()
	if err != nil {
		return err
	}
	p, err := c.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, dest)
}

func SAdd(key string, value string) error {
	c, err := GetRedisConnection()
	if err != nil {
		return err
	}
	return c.SAdd(context.Background(), key, value, 0).Err()
}

func SMembers(key string) ([]string, error) {
	c, err := GetRedisConnection()
	if err != nil {
		return nil, err
	}
	p, err := c.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return p, nil
}
