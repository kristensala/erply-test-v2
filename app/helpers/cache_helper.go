package helpers

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)


func GetFromCache(client redis.Client, key string) string {
    pingErr := client.Ping(context.Background()).Err()
    if pingErr != nil {
        return "" 

    }
    cacheResult := client.Get(context.Background(), key).Val()
    return cacheResult
}

func SetCacheKeyValue(client redis.Client, key string, value interface{}, expirationTimeMinutes time.Duration) {
    pingErr := client.Ping(context.Background()).Err()
    if pingErr != nil {
        return
    }

    client.Set(context.Background(), key, value, expirationTimeMinutes * time.Minute)
}
