package router

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/app/http/request"
)

var AuthRouter = kernel.HTTRouter{
	Route: []kernel.Route{
		{
			Name:        "login",
			Path:        "/login",
			Action:      "AuthController@JWTAuthentication",
			Method:      "POST",
			Validation:  &request.Credentials{},
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
			Validation:  &request.Credentials{},
			Description: "Basic authentication",
			Middleware: []kernel.Middleware{
				middleware.LoggingMiddleware{},
			},
		},
	},
}
