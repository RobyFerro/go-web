package http

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"ikdev/go-web/app/config"
	"ikdev/go-web/controller"
	"ikdev/go-web/exception"
	"ikdev/go-web/helper"
	"net/http"
)

// This method will return an array of models.
// Used to handle migration, seeding and drop operations.
// Every time you add a new model you should register it in this method
func GetControllers(res http.ResponseWriter, req *http.Request) []interface{} {
	var baseController controller.BaseController
	if err := ServiceContainer.Invoke(func(db *gorm.DB, c config.Conf, a *helper.Auth) {
		baseController = controller.BaseController{
			DB:       db,
			Response: res,
			Request:  req,
			Config:   c,
			Auth:     a,
		}
	}); err != nil {
		exception.ProcessError(err)
	}

	// Insert implementation
	// Es: Redis, Elastic, Mongo connections
	checkIntegrations(&baseController)
	return register(baseController)
}

// This method will return an array of controllers.
// Used by Go-Web routing
// Every time you add a new controller you should register it in this method
func register(bc controller.BaseController) []interface{} {
	return []interface{}{
		&controller.UserController{BaseController: bc},
		&controller.AuthController{BaseController: bc},
		&controller.HomeController{BaseController: bc},
		// Here is where you've to register your custom controller
	}
}

// Insert implementation
// Es: Redis, Elastic, Mongo connections
func checkIntegrations(base *controller.BaseController) {

	// If is configured MongoDB will be implemented into service container
	if len(appConf.Mongo.Host) > 0 {
		if err := ServiceContainer.Invoke(func(m *mongo.Database) {
			base.Mongo = m
		}); err != nil {
			exception.ProcessError(err)
		}
	}

	// If is configured Redis will be implemented into service container
	if len(appConf.Redis.Host) > 0 {
		if err := ServiceContainer.Invoke(func(r *redis.Client) {
			base.Redis = r
		}); err != nil {
			exception.ProcessError(err)
		}
	}

	// If is configured ElasticSearch will be implemented into service container
	if len(appConf.Elastic.Hosts) > 0 {
		if err := ServiceContainer.Invoke(func(e *elasticsearch.Client) {
			base.Elastic = e
		}); err != nil {
			exception.ProcessError(err)
		}
	}
}
