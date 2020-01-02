package router

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/config"
	"ikdev/smartcherry/middleware"
)

var db *gorm.DB
var conf config.Conf

func WebRouter(conn *gorm.DB, cgf config.Conf) *mux.Router {
	db = conn
	conf = cgf

	router := mux.NewRouter()
	router.HandleFunc("/", homeAction).Methods("GET")
	router.HandleFunc("/users", newUserAction).Methods("POST")
	router.HandleFunc("/login", loginAction).Methods("POST")

	router.Use(middleware.LoggingMiddleware)

	authRouter := router.PathPrefix("/admin").Subrouter()
	authRouter.HandleFunc("/test", testUserAction)
	authRouter.Use(middleware.AuthMiddleware)

	return router
}
