package main

import (
	"ikdev/go-web/command"
	"os"
)

// Main Go-Web entry point.
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
		case "seed":
			command.RunSeeder()
			break
		case "run:server":
			command.RunServer()
			break
		case "run:queue":
			queueName := os.Args[i+1]
			command.RunQueue(queueName)
			break
		}
	}
}
