package register

import (
	foundation "github.com/RobyFerro/go-web-framework"
	base_register "github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/module"
	"github.com/RobyFerro/go-web/router"
)

// BaseEntities returns a struct that contains Go-Web base entities
func BaseEntities() foundation.BaseEntities {
	return foundation.BaseEntities{
		// Commands configuration represent all Go-Web application conf
		// Every command needs to be registered in the following list
		Commands: console.Commands,
		// Controllers will handle all Go-Web controller
		// Here is where you've to register your custom controller
		// If you create a new controller with Julian it will be auto-registered
		Controllers: base_register.ControllerRegister{
			base_register.ControllerRegisterItem{
				Controller: &controller.UserController{},
				Modules: []base_register.DIModule{
					module.MainModule,
				},
			},
			base_register.ControllerRegisterItem{
				Controller: &controller.AuthController{},
				Modules: []base_register.DIModule{
					module.MainModule,
				},
			},
			base_register.ControllerRegisterItem{
				Controller: &controller.HomeController{},
			},
		},
		// CommandServices represent all console services.
		CommandServices: console.Services,
		// Models will handle all application models
		// Here is where you've to register your custom models
		// If you create a new model with Julian it will be auto-registered
		Models: base_register.ModelRegister{
			model.User{},
		},
		// Router contains all application routes
		Router: []base_register.HTTPRouter{
			router.AppRouter,
			router.AuthRouter,
		},
	}
}
