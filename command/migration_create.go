package command

import (
	"fmt"
	"ikdev/go-web/app"
	"ikdev/go-web/app/config"
	"ikdev/go-web/exception"
	"io/ioutil"
	"time"
)

type MigrationCreate struct {
	Signature   string
	Description string
}

func (c *MigrationCreate) Register() {
	c.Signature = "migration:create <name>"
	c.Description = "Create new migration files"
}

// Create migration files
// This method will create two file UP and DOWN.
// UP: Work only for migrate operation
// DOWN: Work only for rollback operation
func (c *MigrationCreate) Run(kernel *app.HttpKernel, args string) {
	date := time.Now().Unix()
	path := config.GetFilePath("database/migration")

	filenameUp := fmt.Sprintf("%s/%d_%s.up.sql", path, date, args)
	filenameDown := fmt.Sprintf("%s/%d_%s.down.sql", path, date, args)

	fmt.Printf("\nCreating new '%s'...\n", filenameUp)

	if err := ioutil.WriteFile(filenameUp, []byte("/* MIGRATION UP */"), 0755); err != nil {
		exception.ProcessError(err)
	}

	fmt.Printf("Created new up migration: %s\n", filenameUp)
	fmt.Printf("Creating new '%s'...\n", filenameDown)

	if err := ioutil.WriteFile(filenameDown, []byte("/* MIGRATION DOWN */"), 0755); err != nil {
		exception.ProcessError(err)
	}

	fmt.Printf("Created new down migration: %s\n", filenameDown)
	fmt.Printf("Do not forget to register it!")
}
