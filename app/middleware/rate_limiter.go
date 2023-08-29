package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/config"
)

const (
    bucketSize = 5
)

// Token Bucket -> Allow 5 request in 30 seconds
func RateLimiter() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        redis := config.ConnectToRedis();
        client := getUser(ctx)

        if client == ""{
            ctx.AbortWithStatus(http.StatusUnauthorized)
        }

        value := redis.Get(context.Background(), client).Val()
        if value == "" {
            redis.Set(context.Background(), client, 1, time.Second * 30)
        } else {
            newValue := redis.Incr(context.Background(), client).Val()

            if newValue > bucketSize {
                ctx.AbortWithStatus(http.StatusTooManyRequests)
            }
        }
    }
}

func getUser(ctx *gin.Context) string {
    authHeader := ctx.Request.Header["Authorization"]
    authorization := strings.Fields(authHeader[0])
    authValues := strings.Split(authorization[1], ":")

    return authValues[0]
}
