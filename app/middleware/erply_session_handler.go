package middleware

import (
	"net/http"
	"time"

	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/gin-gonic/gin"
)

func ErplySessionHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        sessionKey := ctx.GetString("sessionKey")
        sessionExpiryTime := ctx.GetString("sessionKeyExpireTime")

        unixExpiryTime, _ := time.Parse(time.RFC1123, sessionExpiryTime)
        if sessionKey == "" || sessionExpiryTime == "" || time.Now().After(unixExpiryTime) {
            sessionKey, err := auth.VerifyUser(
                "salakristen1@gmail.com",
                "Qwerty1234",
                "532805",
                http.DefaultClient)

            if err != nil {
                ctx.Error(err)
            }

            sessionInfo, err := auth.GetSessionKeyInfo(sessionKey, "532805", http.DefaultClient)
            if err != nil {
                ctx.Error(err)
            }

            ctx.Set("sessionKey", sessionKey)
            ctx.Set("sessionKeyExpireTime", sessionInfo.ExpireUnixTime)
        }

        ctx.Next()
    }
}
