package http

import (
	"github.com/gorilla/mux"
)

func WebRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeAction).Methods("GET")
	router.HandleFunc("/users", NewUserAction).Methods("POST")
	router.HandleFunc("/login", LoginAction).Methods("POST")
	router.Use(LoggingMiddleware)

	authRouter := router.PathPrefix("/admin").Subrouter()
	authRouter.HandleFunc("/test", TestUserAction)
	authRouter.Use(RefreshTokenMiddleware, AuthMiddleware)

	return router
}
