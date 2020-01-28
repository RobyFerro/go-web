package command

import (
	"go.uber.org/dig"
	"ikdev/go-web/http"
)

type ServerRun struct {
	Signature string
}

// Start Go-Web server
func (c *ServerRun) Run(sc *dig.Container) {
	http.StartServer(sc)
}
