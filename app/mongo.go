package app

import (
	"context"
	"fmt"
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/exception"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

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
