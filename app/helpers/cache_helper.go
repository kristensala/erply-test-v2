package helpers

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetFromCache(client redis.Client, key string) string {
    err := client.Ping(context.Background()).Err()
    if err != nil {
        return "" 
    }
    cacheResult := client.Get(context.Background(), key).Val()
    return cacheResult
}

func SetCacheKeyValue(client redis.Client, key string, value interface{}, expirationTimeMinutes time.Duration) {
    err := client.Ping(context.Background()).Err()
    if err != nil {
        return
    }

    client.Set(context.Background(), key, value, expirationTimeMinutes * time.Minute)
}
