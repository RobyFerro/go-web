package controller

import (
	"fmt"
	"ikdev/go-web/exception"
	"net/http"
)

type HomeController struct {
	BaseController
}

// Main method
func (c *HomeController) Main() {
	c.Response.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(c.Response, "Welcome to smartcherry.io!"); err != nil {
		exception.ProcessError(err)
	}
}
