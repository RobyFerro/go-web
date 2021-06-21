package controller

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/tool"

	"net/http"
)

type HomeController struct {
	kernel.BaseController
}

// Main returns the Go-Web welcome page. This is just an example of Go-Web controller. With the "c" parameter you've access to
// the method/properties declared in BaseController (controller.go).
// Of course you can edit this method with a custom logic.
func (c *HomeController) Main() {
	c.Response.WriteHeader(http.StatusOK)
	tool.View("index.html", c.Response, nil)
}
