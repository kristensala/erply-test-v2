package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/services"
)

type CustomerController interface {
    Add(c *gin.Context)
    Delete(c *gin.Context)
    GetAll(c *gin.Context)
    GetById(c *gin.Context)
    Update(c *gin.Context)
}

type CustomerControllerImpl struct {
    CustomerService services.CustomerService
}

func CustomerControllerInit(customerService services.CustomerService) *CustomerControllerImpl {
    return &CustomerControllerImpl{
        CustomerService: customerService,
    }
}

// Get all customers
// @Summary get all customers details
// @Tags Customer
// @Produce json
// @Success 200 {object} models.ApiResponse{data=[]models.CustomerResponse}
// @Router /customer/getAll [get]
func (c CustomerControllerImpl) GetAll(ctx *gin.Context) {
    c.CustomerService.GetAll(ctx)
}

// Get customer by id
// @Summary get customer by id
// @Tags Customer
// @Produce json
// @Success 200 {object} models.ApiResponse{data=models.CustomerResponse}
// @Router /customer/get/{id} [get]
func (c CustomerControllerImpl) GetById(ctx *gin.Context) {
    customerId := ctx.Param("id")
    c.CustomerService.GetById(ctx, customerId)
}

// Create a new customer
// @Summary create a new customer
// @Tags Customer
// @Produce json
// @Param data body models.CreateCustomerRequest true "customer data"
// @Success 200 {object} models.ApiResponse{data=object}
// @Router /customer/create [post]
func (c CustomerControllerImpl) Add(ctx *gin.Context) {
    c.CustomerService.Add(ctx)
}

// Update an existing customer
// @Summary update an existing customer
// @Tags Customer
// @Produce json
// @Param data body models.UpdateCustomerRequest true "customer data"
// @Success 200 {object} models.ApiResponse{data=object}
// @Router /customer/update [post]
func (c CustomerControllerImpl) Update(ctx *gin.Context) {
    c.CustomerService.Update(ctx)
}

// Delete a customer
// @Summary delete a customer from the system
// @Tags Customer
// @Produce json
// @Param id path string true "customer id"
// @Success 200 {object} models.ApiResponse{data=object}
// @Router /customer/remove/{id} [delete]
func (c CustomerControllerImpl) Delete(ctx *gin.Context) {
    customerId := ctx.Param("id")
    c.CustomerService.Delete(ctx, customerId)
}

