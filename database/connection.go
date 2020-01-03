package database

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
)

func ConnectDB(conf config.Conf) *gorm.DB {
	connectionString, driver := createConnectionString(conf)
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		exception.ProcessError(err)
	}

	return db
}

func ConnectRedis(conf config.Conf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		exception.ProcessError(err)
	}

	return client
}

func createConnectionString(conf config.Conf) (string, string) {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Database,
	), conf.Database.Driver
}
