package command

import (
	"go.uber.org/dig"
	"ikdev/go-web/http"
)

// Start Go-Web server
func RunServer(container *dig.Container) {
	http.StartServer(container)
}
