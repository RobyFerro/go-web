package router

import (
	"ikdev/smartcherry/controller"
	"net/http"
)

func homeAction(w http.ResponseWriter, r *http.Request) {
	homeController := controller.HomeController{
		Controller: controller.Controller{
			DB:       db,
			Response: w,
			Request:  r,
			Config:   conf,
		}}

	homeController.Show()
}

func newUserAction(w http.ResponseWriter, r *http.Request) {
	userController := controller.UserController{Controller: controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
	}}

	userController.Insert()
}

func testUserAction(w http.ResponseWriter, r *http.Request) {
	userController := controller.UserController{Controller: controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
	}}

	userController.Profile()
}

func loginAction(w http.ResponseWriter, r *http.Request) {
	authController := controller.AuthController{Controller: controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
	}}

	authController.Login()
}
