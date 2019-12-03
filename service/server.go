package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"ikdev/smartcherry/config"
	"net/http"
)

func GetHttpServer(router *mux.Router, cfg config.Conf) *http.Server {
	serverString := fmt.Sprintf(":%d", cfg.Server.Port)

	return &http.Server{
		Addr:              serverString,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
}
