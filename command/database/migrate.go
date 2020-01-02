package main

import (
	"github.com/jinzhu/gorm"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"ikdev/go-web/service"
	"reflect"
)

func main() {
	container := service.BuildContainer()

	err := container.Invoke(func(db *gorm.DB) {
		models := database.GetModels()

		migrate(models, db)
		seed(models, db)
	})

	if err != nil {
		exception.ProcessError(err)
	}
}

func migrate(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		v := reflect.ValueOf(m)
		method := v.MethodByName("Migrate")
		method.Call([]reflect.Value{reflect.ValueOf(db)})
	}
}

func seed(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		v := reflect.ValueOf(m)
		method := v.MethodByName("Seed")
		method.Call([]reflect.Value{reflect.ValueOf(db)})
	}
}
