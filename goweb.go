package main

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/console"
	console2 "github.com/RobyFerro/go-web/app/console"
	"github.com/RobyFerro/go-web/app/kernel"
	"github.com/common-nighthawk/go-figure"
	"os"
	"reflect"
)

// Main Go-Web entry point.
// Service container will be passed as parameter for every method
func main() {
	var args string
	pwd, _ := os.Getwd()
	_ = os.Setenv("base_path", pwd)

	// Start HTTP Kernel
	httpKernel := kernel.StartKernel()

	// Print Go-Web container
	printHeader()
	// Merge commands
	mixCommands()

	// New command handler
	cmd := console.Commands[os.Args[1]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	v := reflect.ValueOf(cmd).MethodByName("Run")

	if len(os.Args) == 3 {
		args = os.Args[2]
	}

	v.Call([]reflect.Value{reflect.ValueOf(httpKernel), reflect.ValueOf(args), reflect.ValueOf(console.Commands)})
}

// Print Go-Web CLI header
func printHeader() {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()

	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.eu")
}

// Merge framework commands with yours
func mixCommands() {
	for i, c := range console2.Commands {
		console.Commands[i] = c
	}
}
