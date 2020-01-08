package http

import (
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"ikdev/go-web/service"
	"net/http"
)

// Http package service container.
var Container *dig.Container

// Start HTTP server
func StartServer() {
	router := WebRouter()
	Container = service.BuildContainer(router)
	err := Container.Invoke(func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil {
			exception.ProcessError(err)
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}
}
