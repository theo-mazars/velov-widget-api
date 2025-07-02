package redis

import (
	"log/slog"
	"os"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	status, err := Client.Ping().Result()
	if err != nil {
		slog.Error("Unable to reach Redis server", "error", err)
		panic(err)
	}

	slog.Info("Redis client connected", "PING", status)
}
