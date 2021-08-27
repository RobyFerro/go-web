package register

import (
	"github.com/RobyFerro/go-web-framework"
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
		Middlewares:       Middleware,
	}
}
