package console

import "github.com/RobyFerro/go-web/app/console/command"

// Register all commands here.
// The following map of interfaces expose all available method that can be used by Go-Web CLI tool.
// The map index determines the command that you've to run to for use the relative method.
// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
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
		"install":            &command.Install{},
		// Here is where you've to register your custom controller
	}
)
