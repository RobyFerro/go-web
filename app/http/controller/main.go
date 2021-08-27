package controller

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/tool"
)

type HomeController struct {
	kernel.BaseController
}

// Main returns the Go-Web welcome page. This is just an example of Go-Web controller. With the "c" parameter you've access to
// the method/properties declared in BaseController (controller.go).
// Of course you can edit this method with a custom logic.
func (c *HomeController) Main() {
	tool.View("index.html", c.Response, nil)
}
