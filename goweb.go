package main

import (
	"ikdev/go-web/command"
	"ikdev/go-web/http"
	"ikdev/go-web/service"
	"os"
	"strconv"
)

// Main Go-Web entry point.
// Service container will be passed as parameter for every method
// Todo: Improve this method with reflection
func main() {
	arguments := os.Args
	router := http.WebRouter()
	container := service.BuildContainer(router)

	for i := range arguments {
		if i == 0 {
			continue
		}

		switch os.Args[i] {
		case "migrate:up":
			command.Migrate()
			break
		case "seed":
			command.RunSeeder(container)
			break
		case "run:server":
			command.RunServer(container)
			break
		case "run:daemon":
			command.RunDaemon(container)
			break
		case "run:queue":
			queueName := os.Args[i+1]
			command.RunQueue(queueName, container)
			break
		case "run:failed":
			command.RetryFailedQueue(container)
			break
		case "make:migration":
			queueName := os.Args[i+1]
			command.CreateMigration(queueName)
			break
		case "migrate:rollback":
			steps, _ := strconv.Atoi(os.Args[i+1])
			command.RollbackMigration(steps)
			break
		}
	}
}
