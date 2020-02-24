package main

import (
	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/app"
	"github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/service"
	"os"
)

var (
	// Controllers will handle all Go-Web controller
	Controllers = gwf.ControllerRegister{
		List: []interface{}{
			&controller.UserController{},
			&controller.AuthController{},
			&controller.HomeController{},
		},
		// Here is where you've to register your custom controller
	}
	// Models will handle all application models
	Models = gwf.ModelRegister{
		List: []interface{}{
			model.User{},
			model.FailedJob{},
			// Here is where you've to register your custom models
		},
	}
	// Services will handle all app services
	Services = gwf.ServiceRegister{
		List: []interface{}{
			app.Configuration,
			service.ConnectDB,
			service.ConnectElastic,
			service.ConnectMongo,
			service.ConnectRedis,
			// Here is where you've register your custom service
		},
	}
	// Configuration represent all Go-Web application conf
	Commands = gwf.CommandRegister{
		List: map[string]interface{}{
			"queue:failed": &console.QueueFailed{},
			"queue:run":    &console.QueueRun{},
			// Here is where you've to register your custom commands
		},
	}
)

// Main Go-Web entry point.
// Service container will be passed as parameter for every method
func main() {
	gwf.Start(os.Args[1:], Commands, Controllers, Services, middleware.Middleware{}, Models)
}
