package register

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/router"
	"github.com/RobyFerro/go-web/service"
)

// BaseEntities returns a struct that contains Go-Web base entities
func BaseEntities() foundation.BaseEntities {
	return foundation.BaseEntities{
		Commands: console.Commands,
		Controllers: register.ControllerRegister{
			List: []interface{}{
				&controller.UserController{},
				&controller.AuthController{},
				&controller.HomeController{},
			},
		},
		Services: register.ServiceRegister{
			List: []interface{}{
				service.ConnectDB,
				service.ConnectElastic,
				service.ConnectMongo,
				service.ConnectRedis,
				// Here is where you've register your custom service
			},
		},
		SingletonServices: register.ServiceRegister{
			List: []interface{}{},
		},
		CommandServices: console.Services,
		Models: register.ModelRegister{
			List: []interface{}{
				model.User{},
			},
		},
		Middlewares: register.MiddlewareRegister{
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
		},
		Router: []kernel.HTTRouter{
			router.AppRouter,
			router.AuthRouter,
		},
	}
}
