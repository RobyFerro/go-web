package controller

import (
	"fmt"
	"net/http"
)

type HomeController struct {
	Controller
}

// Main method
func (c *HomeController) Main() {
	c.Response.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(c.Response, "Welcome to smartcherry.io!"); err != nil {
		fmt.Println(err)
	}
}
