package command

// This method will return an array of commands.
// Used by Go-Web routing
// Every time you add a new controller you should register it in this method
func GetCommands() map[string]interface{} {
	return map[string]interface{}{
		"migration:up":       &MigrationUp{Signature: "migration:up"},
		"migration:create":   &MigrationCreate{Signature: "migration:create"},
		"migration:rollback": &MigrateRollback{Signature: "migration:rollback"},
		"queue:failed":       &QueueFailed{Signature: "queue:failed"},
		"queue:run":          &QueueRun{Signature: "queue:run"},
		"seed":               &Seeder{Signature: "seed"},
		"server:daemon":      &ServerDaemon{Signature: "server:daemon"},
		"server:run":         &ServerRun{Signature: "server:run"},
		// Here is where you've to register your custom command
	}
}
