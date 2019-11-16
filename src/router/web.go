package router

import (
	"github.com/gorilla/mux"
	"smartcharry/src/controller"
	"smartcharry/src/middleware"
)

func WebRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.TestHandler)
	router.Use(middleware.LoggingMiddleware)

	return router
}
