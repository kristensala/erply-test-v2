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
    Add(c *gin.Context)
    Update(c *gin.Context, customerId string)
    Delete(c *gin.Context, customerId string)
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

func (c CustomerServiceImpl) Add(ctx *gin.Context) {
    var request models.CreateCustomerRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.Error(err)
        return
    }

    if request.Code != "" && !helpers.IsValidRegCode(request.Code) {
        ctx.IndentedJSON(http.StatusOK, models.ApiResponse{
            IsSuccess: false,
            Data: nil,
            Message: "Invalid reg code",
        })

        return
    }

    ct, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    apiResponse, apiError := helpers.GetErplyClient(ctx).CustomerManager.SaveCustomer(ct, helpers.JsonToMap(request))
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

func (c CustomerServiceImpl) Update(ctx *gin.Context, customerId string) {
    var request models.UpdateCustomerRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.Error(err)
        return
    }

    request.Id = customerId

    if request.Code != "" && helpers.IsValidRegCode(request.Code) {
        ctx.IndentedJSON(http.StatusOK, models.ApiResponse{
            IsSuccess: false,
            Data: nil,
            Message: "Invalid reg code",
        })

        return
    }

    ct, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    apiResponse, apiError := helpers.GetErplyClient(ctx).CustomerManager.SaveCustomer(ct, helpers.JsonToMap(request))
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

func (c CustomerServiceImpl) Delete(ctx *gin.Context, customerId string) {
    filter := map[string]string {
        "customerID": customerId,
    }

    ct, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    apiError := helpers.GetErplyClient(ctx).CustomerManager.DeleteCustomer(ct, filter)
    if apiError != nil {
        ctx.Error(apiError)
        return
    }

    ctx.IndentedJSON(http.StatusOK, models.ApiResponse{
        IsSuccess: true,
        Data: "Customer deleted",
    })

    return
}

func CustomerServiceInit(cache *redis.Client) *CustomerServiceImpl {
    return &CustomerServiceImpl{
        cache: cache,
    }
}
