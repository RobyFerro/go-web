package main

import (
	"ikdev/go-web/command"
	"ikdev/go-web/http"
	"ikdev/go-web/service"
	"os"
)

// Main Go-Web entry point.
// Service container will be passed as parameter for every method
func main() {
	arguments := os.Args
	router := http.WebRouter()
	container := service.BuildContainer(router)

	for i := range arguments {
		if i == 0 {
			continue
		}

		switch os.Args[i] {
		case "migrate":
			command.RunMigration(container)
			break
		case "migrate:new":
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
		}
	}
}
