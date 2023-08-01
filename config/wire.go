// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"

	"github.com/kristensala/erply-test-v2/app/controllers"
	"github.com/kristensala/erply-test-v2/app/services"
)

var cache = wire.NewSet(ConnectToRedis)

var customerServiceSet = wire.NewSet(services.CustomerServiceInit,
    wire.Bind(new(services.CustomerService), new(*services.CustomerServiceImpl)))

var customerControllerSet = wire.NewSet(controllers.CustomerControllerInit,
    wire.Bind(new(controllers.CustomerController), new(*controllers.CustomerControllerImpl)))

func Init() *Initialization {
    wire.Build(NewInitialization, cache, customerServiceSet, customerControllerSet)
    return nil
}

