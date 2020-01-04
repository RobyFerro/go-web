package main

import (
	"ikdev/go-web/command"
	"os"
)

func main() {
	arguments := os.Args

	for i := range arguments {
		if i == 0 {
			continue
		}

		switch os.Args[i] {
		case "migrate":
			command.RunMigration()
			break
		case "run:server":
			command.RunServer()
			break
		case "run:job":
			break
		}
	}
}
