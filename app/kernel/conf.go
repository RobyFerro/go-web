package kernel

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(3)

	routes := config.ConfigurationWeb()
	router := mux.NewRouter()

	go func() {
		handleSingleRoute(routes.Routes, router)
		wg.Done()
	}()

	go func() {
		handleGroups(routes.Groups, router)
		wg.Done()
	}()

	go func() {
		giveAccessToPublicFolder(router)
		wg.Done()
	}()

	wg.Wait()

	return router
}
