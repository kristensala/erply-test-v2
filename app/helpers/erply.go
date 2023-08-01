package helpers

import (
	"net/http"

	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/gin-gonic/gin"
)

func GetErplyClient(ctx *gin.Context) api.Client {
    sessionKey := ctx.MustGet("sessionKey").(string)
    client, err := api.NewClient(sessionKey, "532805", http.DefaultClient)

    if err != nil {
        ctx.Error(err)
    }

    return *client
}
