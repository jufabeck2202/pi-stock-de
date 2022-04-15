package beju

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type repository struct {
	redisClient *redis.Client
}

func (r *repository) NewRedisRepository() (*repository, error) {
	// Define Redis database number.

	// Set Redis options.
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	// Connect to Redis server.
	fmt.Println("Connecting to Redis server...")
	newClient := redis.NewClient(options)

	return &repository{redisClient: newClient}, nil
}

func (r *repository) Set(key string, value interface{}, ttl time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.redisClient.Set(context.Background(), key, p, ttl).Err()
}

func (r *repository) Get(key string, dest interface{}) error {
	p, err := r.redisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, dest)
}

func (r *repository) GetBool(key string) bool {
	data := r.redisClient.Get(context.Background(), key)
	verified, err := data.Bool()
	if err != nil {
		return false
	}
	return verified
}

func (r *repository) SAdd(key string, value string) error {
	return r.redisClient.SAdd(context.Background(), key, value, 0).Err()
}

func (r *repository) SMembers(key string) ([]string, error) {
	p, err := r.redisClient.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return p, nil
}
