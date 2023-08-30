package db

import (
	"X-TENTIONCREW/auth_svc/pkg/config"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func InitRedis(c *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
		Password: "", // No password
		DB:       c.RedisDB,
	})

	// Test the connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Failed to connect to Redis:", err)
	}

	return client, nil
}
