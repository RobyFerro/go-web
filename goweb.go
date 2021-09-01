package main

import (
	"encoding/gob"
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/config"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/register"
	"github.com/joho/godotenv"
	"log"
)

// Main Go-Web entry point.
func main() {
	gob.Register(&model.User{})

	entities := register.BaseEntities()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	foundation.Start(entities, config.GetSever())
}
