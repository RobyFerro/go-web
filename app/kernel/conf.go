package kernel

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"go.uber.org/dig"
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
