package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSW"),
		DB:       0,
	})

	// untuk memastikan koneksi berhasil
	_, err := rdb.Ping(RedisCtx).Result()
	if err != nil {
		log.Fatalf("❌ Failed connection to Redis: %v", err)
	}

	RedisClient = rdb

	log.Println("✅ Redis connection successfully")
}
