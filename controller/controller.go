package controller

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"ikdev/go-web/config"
	"ikdev/go-web/helper"
	"net/http"
)

// Main controller structure
// This structure will be extended by every controllers
type BaseController struct {
	DB       *gorm.DB              // Provide access to MySql instance
	Response http.ResponseWriter   // HTTP response
	Request  *http.Request         // HTTP request
	Config   config.Conf           // Go-Web configuration
	Auth     *helper.Auth          // Authentication/Authorization method
	Redis    *redis.Client         // Provide access to Redis instance
	Mongo    *mongo.Database       // Provide access to MongoDB instance
	Elastic  *elasticsearch.Client // Provide access to ElasticSearch instance
}
