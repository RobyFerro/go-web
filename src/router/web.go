package router

import (
	"github.com/gorilla/mux"
	"smartcharry/src/controller"
)

func WebRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.TestHandler)

	return router
}
