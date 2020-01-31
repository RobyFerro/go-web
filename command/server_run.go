package command

import (
	"ikdev/go-web/app/kernel"
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
func (c *ServerRun) Run(kernel *kernel.HttpKernel, args string, console map[string]interface{}) {
	http.StartServer(kernel.Container)
}
