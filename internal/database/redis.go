package database

import (
	"context"
	"fmt"
	"go-links/configs"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(secret *configs.Secrets) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", secret.Redis.Host, secret.Redis.Port),
		Password: secret.Redis.Password,
		DB:       secret.Redis.Database,
	})

	// Test Connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	return rdb
}
