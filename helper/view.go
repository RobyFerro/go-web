package helper

import (
	"fmt"
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/exception"
	"html/template"
	"net/http"
)

// Render specific view
func View(path string, w http.ResponseWriter, content *interface{}) {
	viewPath := config.GetDynamicPath(fmt.Sprintf("assets/view/%s", path))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles(viewPath)
	if err != nil {
		exception.ProcessError(err)
	}

	_ = t.Execute(w, content)
}
