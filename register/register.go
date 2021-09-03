package register

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/router"
	"github.com/RobyFerro/go-web/service"
)

// BaseEntities returns a struct that contains Go-Web base entities
func BaseEntities() foundation.BaseEntities {
	return foundation.BaseEntities{
		// Commands configuration represent all Go-Web application conf
		// Every command needs to be registered in the following list
		Commands: console.Commands,
		// Controllers will handle all Go-Web controller
		// Here is where you've to register your custom controller
		// If you create a new controller with Alfred it will be auto-registered
		Controllers: register.ControllerRegister{
			&controller.UserController{},
			&controller.AuthController{},
			&controller.HomeController{},
		},
		// Services will handle all app IOC services
		// Every service needs to be registered in the following list
		Services: register.ServiceRegister{
			service.ConnectDB,
			service.ConnectElastic,
			service.ConnectMongo,
			service.ConnectRedis,
		},
		// SingletonServices represent all IOC services that have to be initialized only once.
		// Every service needs to be registered in the following list
		SingletonServices: register.ServiceRegister{},
		// CommandServices represent all console services.
		CommandServices: console.Services,
		// Models will handle all application models
		// Here is where you've to register your custom models
		// If you create a new model with Alfred it will be auto-registered
		Models: register.ModelRegister{
			model.User{},
		},
		// Router contains all application routes
		Router: []register.HTTPRouter{
			router.AppRouter,
			router.AuthRouter,
		},
	}
}
