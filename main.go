package main

import (
	"ikdev/smartcherry/exception"
	"ikdev/smartcherry/service"
	"net/http"
)

func main() {

	container := service.BuildContainer()
	err := container.Invoke(func(server *http.Server) {
		if err := server.ListenAndServe(); err != nil {
			exception.ProcessError(err)
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}
}
