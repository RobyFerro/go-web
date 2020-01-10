package service

import (
	"crypto/tls"
	"fmt"
	"ikdev/go-web/config"
	"net/http"

	"github.com/gorilla/mux"
)

// Prepare HTTP server for Service Container
func GetHttpServer(router *mux.Router, cfg config.Conf) *http.Server {
	serverString := fmt.Sprintf("%s:%d", cfg.Server.Name, cfg.Server.Port)

	var httpServerConf = http.Server{}

	if cfg.Server.Ssl {

		sslCfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		// Add TLS configuration to http server

		httpServerConf = http.Server{
			Addr:    serverString,
			Handler: router,
			//			ReadTimeout:  time.Duration(agentconfig.Ag.Agent.HttpRTimeout) * time.Second,
			//			WriteTimeout: time.Duration(agentconfig.Ag.Agent.HttpWTimeout) * time.Second,
			TLSConfig:    sslCfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}

	} else {
		httpServerConf = http.Server{
			Addr:    serverString,
			Handler: router,
		}
	}

	return &httpServerConf
}
