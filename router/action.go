package router

import (
	"ikdev/go-web/controller"
	"net/http"
)

type Action struct {
	Controller controller.Controller
}

func NewAction(w http.ResponseWriter, r *http.Request) *Action {
	action := new(Action)
	action.Controller = controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
		Auth:     User,
	}

	return action
}

func homeAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	homeController := controller.HomeController{
		Controller: action.Controller,
	}

	homeController.Show()
}

func newUserAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	userController := controller.UserController{
		Controller: action.Controller,
	}

	userController.Insert()
}

func testUserAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	userController := controller.UserController{
		Controller: action.Controller,
	}

	userController.Profile()
}

func loginAction(w http.ResponseWriter, r *http.Request) {
	action := NewAction(w, r)
	authController := controller.AuthController{
		Controller: action.Controller,
	}

	authController.Login()
}
