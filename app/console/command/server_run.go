package command

import (
	"ikdev/go-web/app/http"
	"ikdev/go-web/app/kernel"
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
