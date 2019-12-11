package controller

import (
	"fmt"
	"net/http"
)

type HomeController struct {
	Controller
}

func (c *HomeController) Show() {
	c.Response.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(c.Response, "Welcome to smartcherry.io!"); err != nil {
		fmt.Println(err)
	}
}
