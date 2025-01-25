package initializers

import (
	"log"
	"github.com/go-redis/redis/v8"
	"context"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Redis client initialized successfully!")
}