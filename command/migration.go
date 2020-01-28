package command

import (
	"crypto/sha256"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type Migration struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Hash  string `gorm:"type:varchar(255)"`
	Batch int    `gorm:"type:int(11)"`
}

// Execute MySql migration
// This method will read the contents of database/migrations directory.
func Migrate() {
	db := database.ConnectDB(config.Configuration())
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

// Execute the rollback of the migration
func RollbackMigration(step int) {
	db := database.ConnectDB(config.Configuration())
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

		if payload, err := ioutil.ReadFile(rollbackFile); err != nil {
			exception.ProcessError(err)
		} else {
			db.Exec(string(payload)).Row()
		}

		if err := db.Unscoped().Delete(&m).Error; err != nil {
			exception.ProcessError(err)
		}
	}
}

// Create migration files
// This method will create two file UP and DOWN.
// UP: Work only for migrate operation
// DOWN: Work only for rollback operation
func CreateMigration(name string) {
	date := time.Now().Unix()
	path := config.GetFilePath("database/migration")

	filenameUp := fmt.Sprintf("%s/%d_%s.up.sql", path, date, name)
	filenameDown := fmt.Sprintf("%s/%d_%s.down.sql", path, date, name)

	if err := ioutil.WriteFile(filenameUp, []byte("/* MIGRATION UP */"), 0755); err != nil {
		exception.ProcessError(err)
	}

	if err := ioutil.WriteFile(filenameDown, []byte("/* MIGRATION DOWN */"), 0755); err != nil {
		exception.ProcessError(err)
	}
}

// Run database migrations
// Deprecated: use Migrate method instead
func RunMigration(sc *dig.Container) {
	err := sc.Invoke(func(db *gorm.DB) {
		models := database.GetModels()
		migrate(models, db)
	})

	if err != nil {
		exception.ProcessError(err)
	}
}

// Parse model register and run every migration
// Deprecated: used by deprecated method "RunMigration"
// this method was used by Gorm migration
func migrate(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		v := reflect.ValueOf(m)
		method := v.MethodByName("Migrate")
		method.Call([]reflect.Value{reflect.ValueOf(db)})
	}
}

// Run database seeder
func RunSeeder(sc *dig.Container) {
	err := sc.Invoke(func(db *gorm.DB) {
		models := database.GetModels()
		seed(models, db)
	})

	if err != nil {
		exception.ProcessError(err)
	}
}

// Parse model register and run every seed
func seed(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		v := reflect.ValueOf(m)
		method := v.MethodByName("Seed")
		method.Call([]reflect.Value{reflect.ValueOf(db)})
	}
}

// Set new batch to handle future rollback
func getLastBatch(db *gorm.DB) int {
	m := struct {
		Batch int
	}{}

	db.Raw("SELECT MAX(batch) as batch FROM migrations").Scan(&m)

	return m.Batch
}

// Insert current migration as done into migrations table
func setMigrationAsDone(db *gorm.DB, hash string, name string, batch int) {
	m := Migration{
		Name:  name,
		Hash:  hash,
		Batch: batch,
	}

	db.Create(&m)
}

// Execute migration
func executeMigration(db *gorm.DB, migration string, hash string, batch int) {
	if payload, err := ioutil.ReadFile(migration); err != nil {
		exception.ProcessError(err)
	} else {
		db.Exec(string(payload)).Row()
	}

	setMigrationAsDone(db, hash, migration, batch)
}

// Detect if current migration was already executed
func migrationIsPresent(db *gorm.DB, hash string) bool {
	var count int
	row := db.Table("migrations").Where("hash = ?", hash).Select("count(*)").Row()
	_ = row.Scan(&count)

	return count > 0
}

// Retrieve all migration files located in database/migration folder.
func getAllMigrations() []string {
	var migrations []string
	err := filepath.Walk(config.GetFilePath("database/migration"), func(path string, info os.FileInfo, err error) error {
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
