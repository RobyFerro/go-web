package http

import (
	"ikdev/smartcherry/exception"
	"ikdev/smartcherry/service"
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
