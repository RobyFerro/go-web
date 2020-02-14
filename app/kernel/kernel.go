package kernel

import (
	"encoding/gob"
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web-framework/service"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"go.uber.org/dig"
)

var (
	// Register service container
	Container   *dig.Container
	Controllers = []interface{}{
		&controller.UserController{},
		&controller.AuthController{},
		&controller.HomeController{},
		// Here is where you've to register your custom controller
	}
	Models = []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}
)

// Return an HTTP kernel instance.
func StartKernel() *gwf.HttpKernel {
	Container = service.BuildContainer(Controllers, middleware.Middleware{})

	// Register user model struct to allow easy session encoding/decoding
	// Used only by the basic authentication.
	gob.Register(model.User{})

	return &gwf.HttpKernel{
		Models:    Models,
		Container: Container,
	}
}
