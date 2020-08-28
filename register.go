package main

import (
	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/app"
	"github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/service"
)

var (
	// Controllers will handle all Go-Web controller
	// Here is where you've to register your custom controller
	// If you create a new controller with Alfred it will be auto-registered
	Controllers = gwf.ControllerRegister{
		List: []interface{}{
			&controller.UserController{},
			&controller.AuthController{},
			&controller.HomeController{},
		},
	}
	// Models will handle all application models
	// Here is where you've to register your custom models
	// If you create a new model with Alfred it will be auto-registered
	Models = gwf.ModelRegister{
		List: []interface{}{
			model.User{},
			model.FailedJob{},
		},
	}
	// Services will handle all app services
	// Every service needs to be registered in the following list
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
	// Every command needs to be registered in the following list
	Commands = gwf.CommandRegister{
		List: map[string]interface{}{
			"queue:failed": &console.QueueFailed{},
			"queue:run":    &console.QueueRun{},
			// Here is where you've to register your custom commands
		},
	}
)
