package http

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"ikdev/go-web/config"
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
	if err := container.Invoke(func(db *gorm.DB, c config.Conf, a *helper.Auth) {
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

// App controller register
func register(bc controller.BaseController) []interface{} {
	return []interface{}{
		&controller.UserController{BaseController: bc},
		&controller.AuthController{BaseController: bc},
		&controller.HomeController{BaseController: bc},
	}
}

// Insert implementation
// Es: Redis, Elastic, Mongo connections
func checkIntegrations(base *controller.BaseController) {
	if len(appConf.Mongo.Host) > 0 {
		if err := container.Invoke(func(m *mongo.Database) {
			base.Mongo = m
		}); err != nil {
			exception.ProcessError(err)
		}
	}

	if len(appConf.Redis.Host) > 0 {
		if err := container.Invoke(func(r *redis.Client) {
			base.Redis = r
		}); err != nil {
			exception.ProcessError(err)
		}
	}

	if len(appConf.Elastic.Hosts) > 0 {
		if err := container.Invoke(func(e *elasticsearch.Client) {
			base.Elastic = e
		}); err != nil {
			exception.ProcessError(err)
		}
	}
}
