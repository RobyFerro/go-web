package main

import (
	"encoding/gob"
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/config"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/register"
)

// Main Go-Web entry point.
func main() {
	gob.Register(&model.User{})

	entities := register.BaseEntities()
	foundation.Start(entities, config.ServerConf)
}
