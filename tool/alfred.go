package main

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/register"
	"os"
)

func main() {
	entities := register.BaseEntities()
	foundation.StartCommand(os.Args[1:], entities)
}
