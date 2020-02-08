package kernel

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/gorilla/mux"
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
	// Register all controllers. Every controller must be implemented as a pointer to struct.
	// See the basic implementation to see the correct way to register a new controller.
	Controllers = []interface{}{
		&controller.UserController{},
		&controller.AuthController{},
		&controller.HomeController{},
		// Here is where you've to register your custom controller
	}
	// Register service container
	Container *dig.Container
)

// Parse routing structures and set every route.
// Return a Gorilla Mux router instance with all routes indicated in router.yml file.
func WebRouter() *mux.Router {
	routes := config.ConfigurationWeb()
	router := mux.NewRouter()

	handleSingleRoute(routes.Routes, router)
	handleGroups(routes.Groups, router)
	giveAccessToPublicFolder(router)

	return router
}
