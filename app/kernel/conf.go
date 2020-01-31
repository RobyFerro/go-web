package kernel

import (
	"go.uber.org/dig"
	"ikdev/go-web/app/config"
	"ikdev/go-web/app/http/controller"
	"ikdev/go-web/database/model"
)

var (
	// Declaring base controller
	BaseController controller.BaseController
	// Register web router
	Router = WebRouter()
	// Get app configuration
	Configuration = config.Configuration()
	// Register all models
	Models = []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}
	// Register all controllers
	Controllers = []interface{}{
		&controller.UserController{},
		&controller.AuthController{},
		&controller.HomeController{},
		// Here is where you've to register your custom controller
	}
	// Register service container
	Container *dig.Container
)
