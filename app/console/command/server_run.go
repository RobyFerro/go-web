package command

import (
	"github.com/RobyFerro/go-web-framework"
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
func (c *ServerRun) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	if err := gwf.StartServer(kernel.Container); err != nil {
		gwf.ProcessError(err)
	}
}
