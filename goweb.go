package main

import (
	"encoding/gob"
	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/RobyFerro/go-web/database/model"
	"os"
)

// Main Go-Web entry point.
func main() {
	gob.Register(&model.User{})
	gwf.Start(os.Args[1:], Commands, Controllers, Services, middleware.Middleware{}, Models)
}
