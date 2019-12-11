package router

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/controller"
	"ikdev/smartcherry/middleware"
	"net/http"
)

var db *gorm.DB

func WebRouter(conn *gorm.DB) *mux.Router {
	db = conn

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.Use(middleware.LoggingMiddleware)

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	controller.HomeController(w, r, db)
}
