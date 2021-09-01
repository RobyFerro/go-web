package router

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web/app/http/middleware"
)

var AuthRouter = kernel.HTTRouter{
	Route: []kernel.Route{
		{
			Name:        "login",
			Path:        "/login",
			Action:      "AuthController@JWTAuthentication",
			Method:      "POST",
			Description: "Perform login",
			Middleware: []kernel.Middleware{
				middleware.LoggingMiddleware{},
			},
		},
		{
			Name:        "basic login",
			Path:        "/basic-auth",
			Action:      "AuthController@BasicAuthentication",
			Method:      "POST",
			Description: "Basic authentication",
			Middleware: []kernel.Middleware{
				middleware.LoggingMiddleware{},
			},
		},
	},
}
