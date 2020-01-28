package command

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"reflect"
)

type Seeder struct {
	Signature string
}

func (c *Seeder) Run(sc *dig.Container, args string) {
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
