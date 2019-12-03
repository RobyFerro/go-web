package service

import (
	"go.uber.org/dig"
	"ikdev/smartcherry/config"
	"ikdev/smartcherry/database"
	"ikdev/smartcherry/exception"
	"ikdev/smartcherry/router"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	if err := container.Provide(router.WebRouter); err != nil {
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

	return container
}
