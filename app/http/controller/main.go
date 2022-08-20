package controller

import (
	"github.com/RobyFerro/go-web-framework/kernel"
)

type HomeController struct {
	kernel.BaseController
}

// Main returns the Go-Web welcome page. This is just an example of Go-Web controller. With the "c" parameter you've access to
// the method/properties declared in BaseController (controller.go).
// Of course you can edit this method with a custom logic.
func (c *HomeController) Main() {
	c.Response.WriteHeader(200)
	c.Request.Header.Set("Content-Type", "application/json")

	_, _ = c.Response.Write([]byte(`{"message":"Welcome to Go-Web"}`))
}
