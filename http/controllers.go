package http

import (
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
	var controllers []interface{}
	var baseController controller.BaseController
	if err := container.Invoke(func(db *gorm.DB, c config.Conf, a *helper.Auth, r *redis.Client, m *mongo.Database) {
		baseController = controller.BaseController{
			DB:       db,
			Response: res,
			Request:  req,
			Config:   c,
			Auth:     a,
			Redis:    r,
			Mongo:    m,
		}
	}); err != nil {
		exception.ProcessError(err)
	}

	// Improve configuration
	userController := controller.UserController{
		BaseController: baseController,
	}

	authController := controller.AuthController{
		BaseController: baseController,
	}

	homeController := controller.HomeController{
		BaseController: baseController,
	}
	// End improve configuration

	controllers = append(controllers, &userController)
	controllers = append(controllers, &authController)
	controllers = append(controllers, &homeController)

	return controllers
}
