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

type Action struct {
	Controller controller.Controller
}

func NewAction(w http.ResponseWriter, r *http.Request) *Action {
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

func HomeAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	homeController := controller.HomeController{
		Controller: action.Controller,
	}

	homeController.Show()
}

func NewUserAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	userController := controller.UserController{
		Controller: action.Controller,
	}

	userController.Insert()
}

func TestUserAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	userController := controller.UserController{
		Controller: action.Controller,
	}

	userController.Profile()
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	authController := controller.AuthController{
		Controller: action.Controller,
	}

	authController.Login()
}
