package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/services"
)

type CustomerController interface {
    GetAll(c *gin.Context)
    GetById(c *gin.Context)
    Add(c *gin.Context)
    Update(c *gin.Context)
    Delete(c *gin.Context)
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

func (c CustomerControllerImpl) Add(ctx *gin.Context) {
    c.CustomerService.Add(ctx)
}

func (c CustomerControllerImpl) Update(ctx *gin.Context) {
    customerId := ctx.Param("id")
    c.CustomerService.Update(ctx, customerId)
}

func (c CustomerControllerImpl) Delete(ctx *gin.Context) {
    customerId := ctx.Param("id")
    c.CustomerService.Delete(ctx, customerId)
}

func CustomerControllerInit(customerService services.CustomerService) *CustomerControllerImpl {
    return &CustomerControllerImpl{
        CustomerService: customerService,
    }

}
