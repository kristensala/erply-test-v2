package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/services"
)

type CustomerController interface {
    GetAll(c *gin.Context)
    GetById(c *gin.Context)
}

type CustomerControllerImpl struct {
    CustomerService services.CustomerService
}

func (c CustomerControllerImpl) GetAll(ctx *gin.Context) {
    c.CustomerService.GetAll(ctx)
}

func (c CustomerControllerImpl) GetById(ctx *gin.Context) {
    customerId := ctx.Param("id")
    c.CustomerService.GetById(ctx, customerId)
}

func CustomerControllerInit(customerService services.CustomerService) *CustomerControllerImpl {
    return &CustomerControllerImpl{
        CustomerService: customerService,
    }

}
