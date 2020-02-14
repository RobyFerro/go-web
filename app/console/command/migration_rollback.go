package command

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"strconv"
	"strings"
)

type MigrateRollback struct {
	Signature   string
	Description string
}

func (c *MigrateRollback) Register() {
	c.Signature = "migration:rollback <steps>"
	c.Description = "Rollback migrations"
}

func (c *MigrateRollback) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	var db *gorm.DB
	if err := kernel.Container.Invoke(func(client *gorm.DB) {
		db = client
	}); err != nil {
		gwf.ProcessError(err)
	}

	step, _ := strconv.Atoi(args)
	batch := getLastBatch(db)

	for i := 0; i < step; i++ {
		var migrations []Migration
		if err := db.Order("created_at", true).Where("batch LIKE ?", batch).Find(&migrations).Error; err != nil {
			gwf.ProcessError(err)
		}

		// Execute given rollback
		rollbackMigrations(migrations, db)
		batch--
	}
}

// Core of rollback method.
// This method will parse a given set of migration and run the relative rollback
func rollbackMigrations(migrations []Migration, db *gorm.DB) {
	for _, m := range migrations {
		rollbackFile := strings.ReplaceAll(m.Name, ".up.sql", ".down.sql")
		fmt.Printf("\nRolling back '%s' migration...\n", rollbackFile)

		if payload, err := ioutil.ReadFile(rollbackFile); err != nil {
			gwf.ProcessError(err)
		} else {
			db.Exec(string(payload)).Row()
		}

		if err := db.Unscoped().Delete(&m).Error; err != nil {
			gwf.ProcessError(err)
		}

		fmt.Printf("Success! %s has been rolled back!", rollbackFile)
	}
}
