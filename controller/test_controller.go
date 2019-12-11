package controller

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request, c *gorm.DB) {
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(w, "Hello world"); err != nil {
		fmt.Println(err)
	}
}
