package controller

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/helper"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// Main controller structure. This structure determines what you can find in the controllers instance.
// Adding something else inside this struct will not directly implement the struct. This because is just a part of the
// controller construction. See the "setBaseController" method inside app/kernel/kernel.go
type BaseController struct {
	DB       *gorm.DB              // Provide access to MySql instance
	Response http.ResponseWriter   // HTTP response
	Request  *http.Request         // HTTP request
	Config   config.Conf           // Go-Web configuration
	Auth     *helper.Auth          // Authentication/Authorization method
	Redis    *redis.Client         // Provide access to Redis instance
	Mongo    *mongo.Database       // Provide access to MongoDB instance
	Elastic  *elasticsearch.Client // Provide access to ElasticSearch instance
	Session  *sessions.CookieStore // Provide access to the CookieStore
}
