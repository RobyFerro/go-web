package command

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Execute MySql migration
// This method will read the contents of database/migrations directory.
func Migrate() {
	db := database.ConnectDB(config.Configuration())
	for _, m := range getAllMigrations() {
		if filepath.Ext(m) != ".sql" {
			continue
		}

		payload, err := ioutil.ReadFile(m)
		if err != nil {
			exception.ProcessError(err)
		}

		// Implement regex to get only UP method
		// (?<=-- UP --\s)([\S\s]*)(.*)?(?=\s-- UP --)
		db.Exec(string(payload))

		// Implement regex to get only DOWN method
		// (?<=-- DOWN --\s)([\S\s]*)(.*)?(?=\s-- DOWN --)
	}

}

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
