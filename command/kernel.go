package command

// This method will return an array of commands.
// Used by Go-Web routing
// Every time you add a new controller you should register it in this method
func GetCommands() map[string]interface{} {
	return map[string]interface{}{
		"migration:up":       &MigrationUp{},
		"migration:create":   &MigrationCreate{},
		"migration:rollback": &MigrateRollback{},
		"queue:failed":       &QueueFailed{},
		"queue:run":          &QueueRun{},
		"database:seed":      &Seeder{},
		"server:daemon":      &ServerDaemon{},
		"server:run":         &ServerRun{},
		"controller:create":  &ControllerCreate{},
		"model:create":       &ModelCreate{},
		"show:route":         &ShowRoute{},
		"show:commands":      &ShowCommands{},
		"cmd:create":         &CmdCreate{},
		// Here is where you've to register your custom command
	}
}
