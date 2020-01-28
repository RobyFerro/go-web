package main

import (
	"ikdev/go-web/command"
	"ikdev/go-web/http"
	"ikdev/go-web/service"
	"os"
	"reflect"
)

// Main Go-Web entry point.
// Service container will be passed as parameter for every method
func main() {
	router := http.WebRouter()
	container := service.BuildContainer(router)

	// New command handler
	commands := command.GetCommands()
	cmd := commands[os.Args[1]]
	v := reflect.ValueOf(cmd).MethodByName("Run")

	if len(os.Args) == 2 {
		v.Call([]reflect.Value{reflect.ValueOf(container)})
	} else if len(os.Args) == 3 {
		v.Call([]reflect.Value{reflect.ValueOf(container), reflect.ValueOf(os.Args[2])})
	}
}
