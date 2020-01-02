package router

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"ikdev/go-web/config"
	"ikdev/go-web/helper"
)

var db *gorm.DB
var conf config.Conf
var User *helper.Auth

func WebRouter(conn *gorm.DB, cgf config.Conf, auth *helper.Auth) *mux.Router {
	db = conn
	conf = cgf
	User = auth

	router := mux.NewRouter()
	router.HandleFunc("/", homeAction).Methods("GET")
	router.HandleFunc("/users", newUserAction).Methods("POST")
	router.HandleFunc("/login", loginAction).Methods("POST")
	router.Use(LoggingMiddleware)

	authRouter := router.PathPrefix("/admin").Subrouter()
	authRouter.HandleFunc("/test", testUserAction)
	authRouter.Use(RefreshTokenMiddleware, AuthMiddleware)

	return router
}
