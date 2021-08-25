package register

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/app/http/middleware"
)

// BaseEntities returns a struct that contains Go-Web base entities
func BaseEntities() foundation.BaseEntities {
	return foundation.BaseEntities{
		Commands:          Commands,
		Controllers:       Controllers,
		Services:          Services,
		SingletonServices: SingletonServices,
		CommandServices:   CommandServices,
		Models:            Models,
		Middlewares:       middleware.Middleware{},
	}
}
