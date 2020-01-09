package command

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"ikdev/go-web/http"
	"reflect"
)

// Run database migrations
func RunMigration(sc *dig.Container) {
	err := sc.Invoke(func(db *gorm.DB) {
		models := http.GetModels()
		migrate(models, db)
	})

	if err != nil {
		exception.ProcessError(err)
	}
}

// Run database seeder
func RunSeeder(sc *dig.Container) {
	err := sc.Invoke(func(db *gorm.DB) {
		models := http.GetModels()
		seed(models, db)
	})

	if err != nil {
		exception.ProcessError(err)
	}
}

// Parse model register and run every migration
func migrate(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		v := reflect.ValueOf(m)
		method := v.MethodByName("Migrate")
		method.Call([]reflect.Value{reflect.ValueOf(db)})
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
