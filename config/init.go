package config

import (
	"github.com/kristensala/erply-test-v2/app/controllers"
	"github.com/kristensala/erply-test-v2/app/services"
)

type Initialization struct {
    CustomerService services.CustomerService
    CustomerController controllers.CustomerController
}

func NewInitialization(
    customerService services.CustomerService,
    customerController controllers.CustomerController,
) *Initialization {
    return &Initialization {
        CustomerService: customerService,
        CustomerController: customerController,
    }
}


