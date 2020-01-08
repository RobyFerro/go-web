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

// Actions are used as binder for route and controllers
// Every action must extends Action structure and implement NewAction method.
// See the below code to have an example of implementation
type Action struct {
	Controller controller.Controller
}

// Main action constructor.
// This method should be called by every actions. It will get every dependencies from service container and create an instance
// of base controller.
func (Action) NewAction(w http.ResponseWriter, r *http.Request) *Action {
	action := new(Action)
	if err := Container.Invoke(func(conf config.Conf, auth *helper.Auth) {
		action.Controller.Config = conf
		action.Controller.Auth = auth
		action.Controller.Response = w
		action.Controller.Request = r

		if len(conf.Redis.Host) > 0 {
			if err := Container.Invoke(func(redis *redis.Client) {
				action.Controller.Redis = redis
			}); err != nil {
				exception.ProcessError(err)
			}
		}

		if len(conf.Database.Host) > 0 {
			if err := Container.Invoke(func(db *gorm.DB) {
				action.Controller.DB = db
			}); err != nil {
				exception.ProcessError(err)
			}
		}

		if len(conf.Mongo.Host) > 0 {
			if err := Container.Invoke(func(mongo *mongo.Database) {
				action.Controller.Mongo = mongo
			}); err != nil {
				exception.ProcessError(err)
			}
		}

	}); err != nil {
		exception.ProcessError(err)
	}

	return action
}

// Action that handle default "/" route.
// Feel free to customize it
func (Action) Main(w http.ResponseWriter, r *http.Request) {
	action := Action{}.NewAction(w, r)
	homeController := controller.HomeController{
		Controller: action.Controller,
	}

	homeController.Main()
}

// Action that handle a user registration.
func (Action) NewUser(w http.ResponseWriter, r *http.Request) {
	action := Action{}.NewAction(w, r)
	userController := controller.UserController{
		Controller: action.Controller,
	}

	userController.Insert()
}

// Test method. It's used to check authenticated user.
func (Action) TestUser(w http.ResponseWriter, r *http.Request) {
	action := Action{}.NewAction(w, r)
	userController := controller.UserController{
		Controller: action.Controller,
	}

	userController.Profile()
}

// Handle user login
func (Action) Login(w http.ResponseWriter, r *http.Request) {
	action := Action{}.NewAction(w, r)
	authController := controller.AuthController{
		Controller: action.Controller,
	}

	authController.Login()
}
