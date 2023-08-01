package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/middleware"
	"github.com/kristensala/erply-test-v2/config"
)

func InitRouter(init *config.Initialization) *gin.Engine {
    router := gin.Default()
    router.Use(middleware.ErrorHandler())

    v1 := router.Group("/api/v1")

    protected := v1.Group("")
    protected.Use(middleware.HandleAuthenticate())
    protected.Use(middleware.ErplySessionHandler())

    cust := protected.Group("/customer")
    cust.GET("/getAll", init.CustomerController.GetAll)
    cust.GET("/get/:id", init.CustomerController.GetById)
    cust.POST("/create", init.CustomerController.Add)
    cust.POST("/update/:id", init.CustomerController.Update)
    cust.DELETE("/remove/:id", init.CustomerController.Delete)

    return router
}
