package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/config"
	"ikdev/smartcherry/exception"
)

func ConnectDB(conf config.Conf) *gorm.DB {
	connectionString, driver := createConnectionString(conf)
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		exception.ProcessError(err)
	}

	return db
}

func createConnectionString(conf config.Conf) (string, string) {
	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Database,
	), conf.Database.Driver
}
