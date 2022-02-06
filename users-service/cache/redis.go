package cache

import (
	"github.com/go-redis/redis"
	"os"
)

func InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func CloseRedis(client *redis.Client) error {
	return client.Close()
}
