package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kristensala/erply-test-v2/app/middleware"
	"github.com/kristensala/erply-test-v2/config"
	"github.com/kristensala/erply-test-v2/docs"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(init *config.Initialization) *gin.Engine {
    router := gin.Default()
    router.Use(middleware.ErrorHandler())

    docs.SwaggerInfo.BasePath = "/api/v1"

    v1 := router.Group("/api/v1")

    protected := v1.Group("")
    protected.Use(middleware.HandleAuthenticate())
    protected.Use(middleware.ErplySessionHandler())

    cust := protected.Group("/customer")
    cust.GET("/getAll", init.CustomerController.GetAll)
    cust.GET("/get/:id", init.CustomerController.GetById)
    cust.POST("/create", init.CustomerController.Add)
    cust.POST("/update", init.CustomerController.Update)
    cust.DELETE("/remove/:id", init.CustomerController.Delete)

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler) )

    return router
}
