package main

import (
	"ikdev/smartcherry/database"
	"reflect"
)

func main() {
	models := database.GetModels()

	for _, m := range models {
		v := reflect.ValueOf(m)
		method := v.MethodByName("Migrate")
		method.Call([]reflect.Value{})
	}
}
