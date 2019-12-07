package main

import (
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/database"
	"ikdev/smartcherry/service"
	"reflect"
)

func main() {
	container := service.BuildContainer()

	container.Invoke(func(db *gorm.DB) {
		models := database.GetModels()

		for _, m := range models {
			v := reflect.ValueOf(m)
			method := v.MethodByName("Migrate")
			method.Call([]reflect.Value{reflect.ValueOf(db)})
		}
	})

}
