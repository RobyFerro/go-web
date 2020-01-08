package command

import (
	"github.com/jinzhu/gorm"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"ikdev/go-web/http"
	"ikdev/go-web/service"
	"reflect"
)

// Run database migrations
func RunMigration() {
	container := service.BuildContainer(http.WebRouter())

	err := container.Invoke(func(db *gorm.DB) {
		models := database.GetModels()
		migrate(models, db)
	})

	if err != nil {
		exception.ProcessError(err)
	}
}

// Run database seeder
func RunSeeder() {
	container := service.BuildContainer(http.WebRouter())

	err := container.Invoke(func(db *gorm.DB) {
		models := database.GetModels()
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
