package kernel

import (
	"encoding/gob"

	gwf "github.com/RobyFerro/go-web-framework"
	gwfs "github.com/RobyFerro/go-web-framework/service"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/service"
	"go.uber.org/dig"
)

var (
	// Container represent the global Go-Web service container
	Container *dig.Container
	// Controllers will handle all Go-Web controller
	Controllers = []interface{}{
		&controller.UserController{},
		&controller.AuthController{},
		&controller.HomeController{},
		// Here is where you've to register your custom controller
	}
	// Models will handle all application models
	Models = []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}
	// Services will handle all app services
	Services = []interface{}{
		service.ConnectDB,
		service.ConnectElastic,
		service.ConnectMongo,
		service.ConnectRedis,
		// Here is where you've register your custom service
	}
	// Configuration represent all Go-Web application conf
	Configuration *gwf.Conf
)

// StartKernel will create the HTTP kernel
func StartKernel() *gwf.HttpKernel {
	Container = gwfs.BuildContainer(Controllers, middleware.Middleware{}, Services)
	Container.Invoke(func(conf *gwf.Conf) {
		Configuration = conf
	})

	// Register user model struct to allow easy session encoding/decoding
	// Used only by the basic authentication.
	gob.Register(model.User{})

	return &gwf.HttpKernel{
		Models:    Models,
		Container: Container,
	}
}
