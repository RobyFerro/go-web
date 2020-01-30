package app

import (
	"go.uber.org/dig"
	"ikdev/go-web/command"
	"ikdev/go-web/database/model"
)

type HttpKernel struct {
	Command   map[string]interface{}
	Models    []interface{}
	Container *dig.Container
}

// Start HTTP kernel
func StartKernel() *HttpKernel {
	return &HttpKernel{
		Command:   Command,
		Models:    Models,
		Container: Container,
	}
}

// Register everything go-web needs to run
var (
	// Register all commands
	Command = map[string]interface{}{
		"migration:up":       &command.MigrationUp{},
		"migration:create":   &command.MigrationCreate{},
		"migration:rollback": &command.MigrateRollback{},
		"queue:failed":       &command.QueueFailed{},
		"queue:run":          &command.QueueRun{},
		"database:seed":      &command.Seeder{},
		"server:daemon":      &command.ServerDaemon{},
		"server:run":         &command.ServerRun{},
		"controller:create":  &command.ControllerCreate{},
		"model:create":       &command.ModelCreate{},
		"show:route":         &command.ShowRoute{},
		"show:commands":      &command.ShowCommands{},
		"cmd:create":         &command.CmdCreate{},
		// Here is where you've to register your custom commands
	}

	// Register all models
	Models = []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}

	// Register service container
	Container *dig.Container
)
