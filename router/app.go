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
			Description: "Main access to Go-Web application",
			Middleware: []register.Middleware{
				middleware.RateLimiterMiddleware{},
				middleware.LoggingMiddleware{},
			},
		},
		{
			Name:        "users",
			Path:        "/users",
			Action:      "UserController@Insert",
			Method:      "POST",
			Description: "Insert new user",
			Middleware: []register.Middleware{
				middleware.LoggingMiddleware{},
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
						middleware.LoggingMiddleware{},
						middleware.RateLimiterMiddleware{},
						middleware.RefreshTokenMiddleware{},
					},
				},
			},
			Middleware: []register.Middleware{
				middleware.AuthMiddleware{},
			},
		},
	},
}
