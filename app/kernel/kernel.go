package kernel

import (
	"encoding/gob"
	"github.com/RobyFerro/go-web-framework"
	gwfs "github.com/RobyFerro/go-web-framework/service"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/service"
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
	Services = []interface{}{
		service.ConnectDB,
		service.ConnectElastic,
		service.ConnectMongo,
		service.ConnectRedis,
		// Here is where you've register your custom service
	}
)

// Return an HTTP kernel instance.
func StartKernel() *gwf.HttpKernel {
	Container = gwfs.BuildContainer(Controllers, middleware.Middleware{}, Services)

	// Register user model struct to allow easy session encoding/decoding
	// Used only by the basic authentication.
	gob.Register(model.User{})

	return &gwf.HttpKernel{
		Models:    Models,
		Container: Container,
	}
}
