package helpers

import (
	"net/http"

	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/constants"
)

func GetErplyClient(ctx *gin.Context) api.Client {
    sessionKey := ctx.MustGet("sessionKey").(string)
    client, err := api.NewClient(sessionKey, constants.ErplyClientCode, http.DefaultClient)

    if err != nil {
        ctx.Error(err)
    }

    return *client
}
