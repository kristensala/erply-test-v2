package services

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/helpers"
	"github.com/kristensala/erply-test-v2/app/models"
	"github.com/redis/go-redis/v9"
)

type CustomerService interface {
    GetAll(c *gin.Context)
    GetById(c *gin.Context, customerId string)
}

type CustomerServiceImpl struct {
    cache *redis.Client
}

func (c CustomerServiceImpl) GetAll(ctx *gin.Context) {
    ct, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    apiResponse, apiError := helpers.GetErplyClient(ctx).CustomerManager.GetCustomers(ct, nil)
    if apiError != nil {
        ctx.Error(apiError)
        return
    }

    ctx.IndentedJSON(http.StatusOK, models.ApiResponse{
        IsSuccess: true,
        Data: apiResponse,
    })

    return
}

func (c CustomerServiceImpl) GetById(ctx *gin.Context, customerId string) {
    filter := map[string]string {
        "customerID": customerId,
    }

    ct, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    apiResponse, apiError := helpers.GetErplyClient(ctx).CustomerManager.GetCustomers(ct, filter)
    if apiError != nil {
        ctx.Error(apiError)
        return
    }

    ctx.IndentedJSON(http.StatusOK, models.ApiResponse{
        IsSuccess: true,
        Data: apiResponse,
    })

    return
}

func CustomerServiceInit(cache *redis.Client) *CustomerServiceImpl {
    return &CustomerServiceImpl{
        cache: cache,
    }
}
