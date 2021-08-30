package register

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/config"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/service"
)

var (
	// Controllers will handle all Go-Web controller
	// Here is where you've to register your custom controller
	// If you create a new controller with Alfred it will be auto-registered
	Controllers = register.ControllerRegister{
		List: []interface{}{
			&controller.UserController{},
			&controller.AuthController{},
			&controller.HomeController{},
		},
	}
	// Models will handle all application models
	// Here is where you've to register your custom models
	// If you create a new model with Alfred it will be auto-registered
	Models = register.ModelRegister{
		List: []interface{}{
			model.User{},
		},
	}
	// Services will handle all app IOC services
	// Every service needs to be registered in the following list
	Services = register.ServiceRegister{
		List: []interface{}{
			config.GetConfiguration,
			service.ConnectDB,
			service.ConnectElastic,
			service.ConnectMongo,
			service.ConnectRedis,
			// Here is where you've register your custom service
		},
	}
	// SingletonServices represent all IOC services that have to be initialized only once.
	// Every service needs to be registered in the following list
	SingletonServices = register.ServiceRegister{
		List: []interface{}{},
	}
	// CommandServices represent all console services.
	CommandServices = console.Services
	// Commands configuration represent all Go-Web application conf
	// Every command needs to be registered in the following list
	Commands = console.Commands
	// Middleware is used to register all application middleware.
	Middleware = register.MiddlewareRegister{
		List: []interface{}{
			&middleware.AuthMiddleware{
				Name:        "Auth",
				Description: "Provides JWT authentication",
			},
			&middleware.BasicAuthMiddleware{
				Name:        "BasicAuth",
				Description: "Provides basic authentication",
			},
			&middleware.LoggingMiddleware{
				Name:        "Logging",
				Description: "Logs every request in console",
			},
			&middleware.RateLimiterMiddleware{
				Name:        "RateLimiter",
				Description: "Provides rate limit over HTTP requests",
			},
			&middleware.RefreshTokenMiddleware{
				Name:        "RefreshToken",
				Description: "Refresh JWT token",
			},
		},
	}
)
