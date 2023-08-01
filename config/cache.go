package config

import (
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "test123",
        DB: 0,
    })

    return client
}
