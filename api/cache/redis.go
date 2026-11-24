package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func initRedis() {
	// REDIS_URL format: redis://user:password@host:6379/0
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		// Don't print err: it can echo the URL, which contains the password.
		fmt.Println("❌ Invalid REDIS_URL")
		os.Exit(1)
	}
	Client = redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		fmt.Println("❌ Error connecting to Redis")
		panic(err)
	}

	fmt.Println("✅ Redis client initialized")
}
