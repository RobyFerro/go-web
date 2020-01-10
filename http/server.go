package http

import (
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
	"net/http"

	"go.uber.org/dig"
)

var container *dig.Container
var appConf config.Conf

// Start HTTP server
func StartServer(sc *dig.Container) {
	container = sc
	err := sc.Invoke(func(s *http.Server, conf config.Conf) {
		// Declaring global app configuration
		appConf = conf

		if appConf.Server.Ssl {

			if httpErr := s.ListenAndServeTLS(appConf.Server.SslCert, appConf.Server.SslKey); httpErr != nil {
				exception.ProcessError(httpErr)
			}

		} else {

			if httpsErr := s.ListenAndServe(); httpsErr != nil {
				exception.ProcessError(httpsErr)
			}

		}
	})

	if err != nil {
		exception.ProcessError(err)
	}
}
