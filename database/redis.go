package database

import (
	"fmt"
	"github.com/RianIhsan/go-postgres-redis/config"
	"github.com/redis/go-redis/v9"
)

func RedisConnection(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	fmt.Println("Redis connection successfully")
	return rdb
}
