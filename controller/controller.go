package controller

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"ikdev/go-web/config"
	"ikdev/go-web/helper"
	"net/http"
)

// Main controller structure
// This structure will be extended by every controllers
type Controller struct {
	DB       *gorm.DB
	Response http.ResponseWriter
	Request  *http.Request
	Config   config.Conf
	Auth     *helper.Auth
	Redis    *redis.Client
	Mongo    *mongo.Database
}
