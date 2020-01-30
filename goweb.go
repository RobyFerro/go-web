package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"ikdev/go-web/app"
	"ikdev/go-web/command"
	"os"
	"reflect"
)

// Main Go-Web entry point.
// Service container will be passed as parameter for every method
func main() {
	var args string
	httpKernel := app.StartKernel()

	// Print Go-Web container
	printHeader()

	// New command handler
	cmd := command.Register[os.Args[1]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	v := reflect.ValueOf(cmd).MethodByName("Run")

	if len(os.Args) == 3 {
		args = os.Args[2]
	}

	v.Call([]reflect.Value{reflect.ValueOf(httpKernel), reflect.ValueOf(args)})
}

// Print Go-Web CLI header
func printHeader() {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()

	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.eu")
}
