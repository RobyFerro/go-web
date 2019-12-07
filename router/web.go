package router

import (
	"github.com/gorilla/mux"
	"ikdev/smartcherry/controller"
	"ikdev/smartcherry/middleware"
)

func WebRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.TestHandler).Methods("GET")
	router.Use(middleware.LoggingMiddleware)

	return router
}
