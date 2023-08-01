package controllers

import "github.com/gin-gonic/gin"

type CustomerController interface {
    GetAll(c *gin.Context)
}

type CustomerControllerImpl struct {

}
