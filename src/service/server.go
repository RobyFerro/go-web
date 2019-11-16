package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"smartcharry/config"
)

func StartHttpServer(router *mux.Router) {
	var cfg config.Conf
	config.GetConf(&cfg)

	serverString := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := http.ListenAndServe(serverString, router); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
