package command

import (
	"crypto/sha256"
	"fmt"
	"github.com/jinzhu/gorm"
	"ikdev/go-web/app/config"
	"ikdev/go-web/app/kernel"
	"ikdev/go-web/exception"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type MigrationUp struct {
	Signature   string
	Description string
}

func (c *MigrationUp) Register() {
	c.Signature = "migration:up"
	c.Description = "Execute migration"
}

type Migration struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Hash  string `gorm:"type:varchar(255)"`
	Batch int    `gorm:"type:int(11)"`
}

func (c *MigrationUp) Run(kernel *kernel.HttpKernel, args string, console map[string]interface{}) {

	var db *gorm.DB
	if err := kernel.Container.Invoke(func(client *gorm.DB) {
		db = client
	}); err != nil {
		exception.ProcessError(err)
	}

	db.AutoMigrate(&Migration{})
	batch := getLastBatch(db) + 1

	for _, m := range getAllMigrations() {
		if !strings.Contains(m, "up.sql") {
			continue
		}

		// Detect if migration was already executed
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(m)))
		if migrationIsPresent(db, hash) {
			continue
		}

		// Exec migration
		executeMigration(db, m, hash, batch)
	}
}

// Retrieve all migration files located in database/migration folder.
func getAllMigrations() []string {
	var migrations []string
	err := filepath.Walk(config.GetDynamicPath("database/migration"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		migrations = append(migrations, path)
		return nil
	})

	if err != nil {
		exception.ProcessError(err)
	}

	return migrations
}

// Set new batch to handle future rollback
func getLastBatch(db *gorm.DB) int {
	m := struct {
		Batch int
	}{}

	db.Raw("SELECT MAX(batch) as batch FROM migrations").Scan(&m)

	return m.Batch
}

// Detect if current migration was already executed
func migrationIsPresent(db *gorm.DB, hash string) bool {
	var count int
	row := db.Table("migrations").Where("hash = ?", hash).Select("count(*)").Row()
	_ = row.Scan(&count)

	return count > 0
}

// Execute migration
func executeMigration(db *gorm.DB, migration string, hash string, batch int) {
	fmt.Printf("\nMigrating '%s'\n", migration)
	if payload, err := ioutil.ReadFile(migration); err != nil {
		exception.ProcessError(err)
	} else {
		db.Exec(string(payload)).Row()
	}

	setMigrationAsDone(db, hash, migration, batch)
}

// Insert current migration as done into migrations table
func setMigrationAsDone(db *gorm.DB, hash string, name string, batch int) {
	m := Migration{
		Name:  name,
		Hash:  hash,
		Batch: batch,
	}

	db.Create(&m)
	fmt.Printf("Success! Migration %s has been executed", name)
}
