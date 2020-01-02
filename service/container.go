package service

import (
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
)

func BuildContainer(router *mux.Router) *dig.Container {
	container := dig.New()

	err := container.Provide(func() *mux.Router {
		return router
	})

	if err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(config.Configuration); err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(database.ConnectDB); err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(GetHttpServer); err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(SetAuth); err != nil {
		exception.ProcessError(err)
	}

	return container
}
