package command

import (
	"go.uber.org/dig"
	"ikdev/go-web/http"
)

type ServerRun struct {
	Signature   string
	Description string
}

func (c *ServerRun) Register() {
	c.Signature = "server:run"
	c.Description = "Run Go-Web server"
}

// Start Go-Web server
func (c *ServerRun) Run(sc *dig.Container, args string) {
	http.StartServer(sc)
}
