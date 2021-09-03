package router

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app/http/middleware"
)

var AppRouter = register.HTTPRouter{
	Route: []register.Route{
		{
			Name:        "home",
			Path:        "/",
			Action:      "HomeController@Main",
			Method:      "GET",
			Description: "Main route",
			Middleware: []register.Middleware{
				middleware.NewRateLimiterMiddleware(),
				middleware.NewLoggingMiddleware(),
			},
		},
		{
			Name:        "users",
			Path:        "/users",
			Action:      "UserController@Insert",
			Method:      "POST",
			Description: "Insert new user",
			Middleware: []register.Middleware{
				middleware.NewLoggingMiddleware(),
			},
		},
	},
	Groups: []register.Group{
		{
			Name:   "admin",
			Prefix: "/admin",
			Routes: []register.Route{
				{
					Name:        "test",
					Path:        "/test",
					Action:      "UserController@Profile",
					Method:      "GET",
					Description: "Test user authentication",
					Middleware: []register.Middleware{
						middleware.NewLoggingMiddleware(),
						middleware.NewRateLimiterMiddleware(),
						middleware.NewRefreshTokenMiddleware(),
					},
				},
			},
			Middleware: []register.Middleware{
				middleware.NewAuthMiddleware(),
			},
		},
	},
}
