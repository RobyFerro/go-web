package router

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web/app/http/middleware"
)

var AppRouter = kernel.HTTRouter{
	Route: []kernel.Route{
		{
			Name:        "home",
			Path:        "/",
			Action:      "HomeController@Main",
			Method:      "GET",
			Description: "Main access to Go-Web application",
			Middleware: []kernel.Middleware{
				&middleware.RateLimiterMiddleware{},
				&middleware.LoggingMiddleware{},
			},
		},
		{
			Name:        "users",
			Path:        "/users",
			Action:      "UserController@Insert",
			Method:      "POST",
			Description: "Insert new user",
			Middleware: []kernel.Middleware{
				&middleware.LoggingMiddleware{},
			},
		},
	},
	Groups: []kernel.Group{
		{
			Name:   "admin",
			Prefix: "/admin",
			Routes: []kernel.Route{
				{
					Name:        "test",
					Path:        "/test",
					Action:      "UserController@Profile",
					Method:      "GET",
					Description: "Test user authentication",
					Middleware: []kernel.Middleware{
						&middleware.LoggingMiddleware{},
						&middleware.RateLimiterMiddleware{},
						&middleware.RefreshTokenMiddleware{},
					},
				},
			},
			Middleware: []kernel.Middleware{
				&middleware.AuthMiddleware{},
			},
		},
	},
}
