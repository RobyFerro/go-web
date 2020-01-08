package http

import (
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"net/http"
)

var container *dig.Container

// Start HTTP server
func StartServer(sc *dig.Container) {
	container = sc
	err := sc.Invoke(func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil {
			exception.ProcessError(err)
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}
}
