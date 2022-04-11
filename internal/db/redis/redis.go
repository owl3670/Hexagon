package redis

import (
	"Hexagon/config"
	"github.com/go-redis/redis/v8"
)

func OpenDB(cfg config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return client
}
