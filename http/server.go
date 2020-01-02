package http

import (
	"ikdev/go-web/exception"
	"ikdev/go-web/service"
	"net/http"
)

func StartServer() {
	container := service.BuildContainer()
	err := container.Invoke(func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil {
			exception.ProcessError(err)
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}
}
