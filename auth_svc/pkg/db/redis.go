package db

import (
	"auth_svc/pkg/config"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func InitRedis(c *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddress, // Use the new RedisAddress field
		Password: "",             // No password
		DB:       0,              // Assuming Redis DB is not used, use your appropriate DB number if applicable
	})

	// Test the connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Failed to connect to Redis:", err)
	}

	return client, nil
}
