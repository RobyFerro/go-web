package http

import (
	"go.uber.org/dig"
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
	"net/http"
)

var container *dig.Container
var appConf config.Conf

// Start HTTP server
func StartServer(sc *dig.Container) {
	container = sc
	err := sc.Invoke(func(s *http.Server, conf config.Conf) {
		// Declaring global app configuration
		appConf = conf

		if err := s.ListenAndServe(); err != nil {
			exception.ProcessError(err)
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}
}
