package console

import "ikdev/go-web/app/console/command"

// This method will return an array of commands.
// Used by Go-Web routing
// Every time you add a new controller you should register it in this method

var (
	Commands = map[string]interface{}{
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
		"middleware:create":  &command.MiddlewareCreate{},
		"job:create":         &command.JobCreate{},
		// Here is where you've to register your custom controller
	}
)
