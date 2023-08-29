package middleware

import (
	"net/http"
	"time"

	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/constants"
)

func ErplySessionHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        sessionKey := ctx.GetString("sessionKey")
        sessionExpiryTime := ctx.GetString("sessionKeyExpireTime")

        unixExpiryTime, err := time.Parse(time.RFC1123, sessionExpiryTime)
        if err != nil {
            ctx.Error(err)
        }

        if sessionKey == "" || sessionExpiryTime == "" || time.Now().After(unixExpiryTime) {
            sessionKey, err := auth.VerifyUser(
                constants.ErplyUsername,
                constants.ErplyPassword,
                constants.ErplyClientCode,
                http.DefaultClient)

            if err != nil {
                ctx.Error(err)
            }

            sessionInfo, err := auth.GetSessionKeyInfo(sessionKey, constants.ErplyClientCode, http.DefaultClient)
            if err != nil {
                ctx.Error(err)
            }

            ctx.Set("sessionKey", sessionKey)
            ctx.Set("sessionKeyExpireTime", sessionInfo.ExpireUnixTime)
        }

        ctx.Next()
    }
}
