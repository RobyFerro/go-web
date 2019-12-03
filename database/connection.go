package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/config"
	"ikdev/smartcherry/exception"
)

func ConnectDB() *gorm.DB {
	connectionString, driver := createConnectionString()
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		exception.ProcessError(err)
	}

	return db
}

func createConnectionString() (string, string) {
	var cfg config.Conf
	config.GetConf(&cfg)

	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Database,
	), cfg.Database.Driver
}
