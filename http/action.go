package http

import (
	"github.com/jinzhu/gorm"
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

	err := Container.Invoke(func(db *gorm.DB, conf config.Conf, auth *helper.Auth) {
		action.Controller.DB = db
		action.Controller.Config = conf
		action.Controller.Auth = auth
		action.Controller.Response = w
		action.Controller.Request = r
	})

	if err != nil {
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
