package app

import (
	"go.uber.org/dig"
	"ikdev/go-web/database/model"
	"ikdev/go-web/http"
	"ikdev/go-web/service"
)

type HttpKernel struct {
	Models    []interface{}
	Container *dig.Container
}

// Start HTTP kernel
func StartKernel() *HttpKernel {
	router := http.WebRouter()
	Container = service.BuildContainer(router)

	return &HttpKernel{
		Models:    Models,
		Container: Container,
	}
}

// Register everything go-web needs to run
var (
	// Register all models
	Models = []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}

	// Register service container
	Container *dig.Container
)
