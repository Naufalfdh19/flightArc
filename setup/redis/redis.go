package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", 
		DB:       0,  
	})

	return rdb
}
