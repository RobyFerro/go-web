package service

import (
	"context"
	"fmt"
	"github.com/RobyFerro/go-web/config"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ConnectMongo returns a connection to MongoDB
func ConnectMongo() *mongo.Database {
	conn := fmt.Sprintf("mongodb://%s:%d", config.MailConf.Host, config.MailConf.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(conn))

	if err != nil {
		log.Error(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	return client.Database(config.MongoConf.Database)
}
