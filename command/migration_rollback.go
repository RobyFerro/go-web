package command

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"io/ioutil"
	"strconv"
	"strings"
)

type MigrateRollback struct {
	Signature string
}

func (c *MigrateRollback) Run(sc *dig.Container, args string) {
	var db *gorm.DB
	if err := sc.Invoke(func(client *gorm.DB) {
		db = client
	}); err != nil {
		exception.ProcessError(err)
	}

	step, _ := strconv.Atoi(args)
	batch := getLastBatch(db)

	for i := 0; i < step; i++ {
		var migrations []Migration
		if err := db.Order("created_at", true).Where("batch LIKE ?", batch).Find(&migrations).Error; err != nil {
			exception.ProcessError(err)
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
			exception.ProcessError(err)
		} else {
			db.Exec(string(payload)).Row()
		}

		if err := db.Unscoped().Delete(&m).Error; err != nil {
			exception.ProcessError(err)
		}

		fmt.Printf("Success! %s has been rolled back!", rollbackFile)
	}
}
