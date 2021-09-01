package main

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/register"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	entities := register.BaseEntities()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	foundation.StartCommand(os.Args[1:], entities)
}
