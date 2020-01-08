package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
	"time"
)

// Connect to sql database
func ConnectDB(conf config.Conf) *gorm.DB {
	connectionString, driver := createConnectionString(conf)
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		exception.ProcessError(err)
	}

	return db
}

// Connect to Redis
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

// Connect to MongoDB
func ConnectMongo(conf config.Conf) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", conf.Mongo.Host, conf.Mongo.Port)))

	if err != nil {
		exception.ProcessError(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	return client.Database(conf.Mongo.Database)
}

// Create string for SQL connection
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
