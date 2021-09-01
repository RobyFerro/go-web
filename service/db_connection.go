package service

import (
	"fmt"
	"github.com/RobyFerro/go-web/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

// ConnectDB to sql database
func ConnectDB() *gorm.DB {
	connectionString, driver := createConnectionString()
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		log.Error(err)
	}

	return db
}

// Create string for SQL connection
func createConnectionString() (string, string) {
	conf := config.GetSQL()
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	), conf.Driver
}
