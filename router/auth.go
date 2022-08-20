package router

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app/http/validation"
)

var AuthRouter = register.HTTPRouter{
	Route: []register.Route{
		{
			Name:        "login",
			Path:        "/login",
			Action:      "AuthController@JWTAuthentication",
			Method:      "POST",
			Validation:  &validation.Credentials{},
			Description: "Perform login",
			Middleware:  []register.Middleware{},
		},
	},
}
