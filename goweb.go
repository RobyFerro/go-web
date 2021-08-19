package main

import (
	"encoding/gob"
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"os"
)

// Main Go-Web entry point.
func main() {
	gob.Register(&model.User{})

	entities := baseEntities()
	foundation.Start(os.Args[1:], entities)
}

// Returns a filled BaseEntities struct.
func baseEntities() foundation.BaseEntities {
	return foundation.BaseEntities{
		Commands:          Commands,
		Controllers:       Controllers,
		Services:          Services,
		SingletonServices: SingletonServices,
		Models:            Models,
		Middlewares:       middleware.Middleware{},
	}
}
