package cache

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping().Result()

	if err != nil {
		fmt.Println("❌ Error connecting to Redis")
		panic(err)
	}

	fmt.Println("✅ Redis client initialized")
}
