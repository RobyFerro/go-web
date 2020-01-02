package http

import (
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"ikdev/go-web/service"
	"net/http"
)

var Container *dig.Container

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
