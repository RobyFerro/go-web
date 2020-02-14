package console

// Register all commands here.
// The following map of interfaces expose all available method that can be used by Go-Web CLI tool.
// The map index determines the command that you've to run to for use the relative method.
// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
var (
	Commands = map[string]interface{}{
		// Here is where you've to register your custom controller
	}
)
