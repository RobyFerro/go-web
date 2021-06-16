package main

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/cli"
	"github.com/RobyFerro/go-web/app"
	"github.com/RobyFerro/go-web/service"
	"github.com/common-nighthawk/go-figure"
	"log"
	"os"
)

func main() {
	myFigure := figure.NewFigure("GWF-Alfred", "", true)
	myFigure.Print()

	fmt.Println("\nGo-Web - Alfred CLI tool - Author: roberto.ferro@ikdev.eu")

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Missing arguments, type -h/--help to show available commands")
	}

	switch args[0] {
	case "--help", "-h":
		help()
	case "--make-controller", "-mC":
		if len(args) != 2 {
			log.Fatal("Missing controller name")
		}
		cmd := cli.ControllerCreate{Args: args[1]}
		cmd.Run()
	case "--make-command", "-mCMD":
		if len(args) != 2 {
			log.Fatal("Missing command name")
		}

		cmd := cli.CmdCreate{Args: args[1]}
		cmd.Run()
	case "--make-model", "-mM":
		if len(args) != 2 {
			log.Fatal("Missing model name")
		}
		cmd := cli.ModelCreate{Args: args[1]}
		cmd.Run()
	case "--make-migration", "-mDBM":
		if len(args) != 2 {
			log.Fatal("Missing migration name")
		}
		cmd := cli.MigrationCreate{Args: args[1]}
		cmd.Run()
	case "--make-job", "-mJ":
		if len(args) != 2 {
			log.Fatal("Missing migration name")
		}
		cmd := cli.JobCreate{Args: args[1]}
		cmd.Run()
	case "--migrate-up", "-mU":
		conf, err := app.Configuration()
		if err != nil {
			log.Fatal(err)
		}

		db := service.ConnectDB(conf)
		cmd := cli.MigrationUp{}
		cmd.Run(db)
	case "--migrate-rollback", "-mR":
		if len(args) != 2 {
			log.Fatal("Missing rollback step")
		}

		conf, err := app.Configuration()
		if err != nil {
			log.Fatal(err)
		}

		db := service.ConnectDB(conf)
		cmd := cli.MigrateRollback{
			Args: args[1],
		}
		cmd.Run(db)
	case "--make-middleware", "-mMW":
		if len(args) != 2 {
			log.Fatal("Missing middleware name")
		}
		cmd := cli.MiddlewareCreate{Args: args[1]}
		cmd.Run()
	case "--generate-key", "-gK":
		cmd := cli.GenerateKey{}
		cmd.Run()
	}
}

func help() {
	fmt.Println("Usage: alfred [command] [options]")

	fmt.Println("\nBASIC COMMANDS:")
	fmt.Println(" --help -h: Shows help menu")

	fmt.Println("\nPROJECT COMMANDS - RUNS THIS COMMANDS ONLY IN PROJECT ROOT!")
	fmt.Println(" --make-controller, -mC <controller_name>: Creates new Go-Web controller")
	fmt.Println(" --make-command, -mCMD <command_name>: Creates new Go-Web command")
	fmt.Println(" --make-model, -mM <model_name>: Creates new Go-Web model")
	fmt.Println(" --make-migration , -mDBM <migration_name>: Creates new Go-Web SQL migration")
	fmt.Println(" --make-middleware , -mMW <middleware_name>: Creates new middleware")
	fmt.Println(" --make-job, -mJ <job_name>: Creates new async job")
	fmt.Println(" --migrate-up, -mU: Executes migrations")
	fmt.Println(" --migrate-rollback, -mR <steps>: Executes migrations rollback")
	fmt.Println(" --generate-key, -gK: Generate new application key")
}
