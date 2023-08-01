package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/models"
)

func ErrorHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Next()

        for _, err := range ctx.Errors {
            ctx.AbortWithStatusJSON(http.StatusInternalServerError, 
                models.ApiResponse{
                    IsSuccess: false,
                    Message: err.Error(),
                })
        }
    }
}
