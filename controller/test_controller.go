package controller

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(w, "Hello world"); err != nil {
		fmt.Println(err)
	}
}
