package service

import (
	"fmt"
	"github.com/RobyFerro/go-web/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

// ConnectDB to sql database
func ConnectDB(conf *app.Conf) *gorm.DB {
	connectionString, driver := createConnectionString(conf)
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		log.Error(err)
	}

	return db
}

// Create string for SQL connection
func createConnectionString(conf *app.Conf) (string, string) {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Database,
	), conf.Database.Driver
}
