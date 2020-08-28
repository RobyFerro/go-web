package main

import (
	"fmt"
	gwf "github.com/RobyFerro/go-web-framework"
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
	case "--create-service", "-cS":
		if len(args) != 2 {
			log.Fatal("Missing destination path")
		}
		cmd := gwf.Install{Args: args[1]}
		cmd.Run()
	case "--make-controller", "-mC":
		if len(args) != 2 {
			log.Fatal("Missing controller name")
		}
		cmd := gwf.ControllerCreate{Args: args[1]}
		cmd.Run()
	case "--make-command", "-mCMD":
		if len(args) != 2 {
			log.Fatal("Missing command name")
		}

		cmd := gwf.CmdCreate{Args: args[1]}
		cmd.Run()
	case "--make-model", "-mM":
		if len(args) != 2 {
			log.Fatal("Missing model name")
		}
		cmd := gwf.ModelCreate{Args: args[1]}
		cmd.Run()
	case "--make-migration", "-mDBM":
		if len(args) != 2 {
			log.Fatal("Missing migration name")
		}
		cmd := gwf.MigrationCreate{Args: args[1]}
		cmd.Run()
	case "--make-job", "-mJ":
		if len(args) != 2 {
			log.Fatal("Missing migration name")
		}
		cmd := gwf.JobCreate{Args: args[1]}
		cmd.Run()
	case "--show-routes", "-sR":
		cmd := gwf.ShowRoute{}
		cmd.Run()
	case "--migrate-up", "-mU":
		conf, err := app.Configuration()
		if err != nil {
			log.Fatal(err)
		}

		db := service.ConnectDB(conf)
		cmd := gwf.MigrationUp{}
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
		cmd := gwf.MigrateRollback{
			Args: args[1],
		}
		cmd.Run(db)
	case "--make-middleware", "-mMW":
		if len(args) != 2 {
			log.Fatal("Missing middleware name")
		}
		cmd := gwf.MiddlewareCreate{Args: args[1]}
		cmd.Run()
	case "--generate-key", "-gK":
		cmd := gwf.GenerateKey{}
		cmd.Run()
	case "--http-load", "-hL":
		if len(args) != 2 {
			log.Fatal("Missing configuration path")
		}

		conf, err := gwf.Configuration()
		if err != nil {
			log.Fatal(err)
		}

		cmd := gwf.HttpLoad{}
		cmd.Run(conf)
	}
}

func help() {
	fmt.Println("Usage: alfred [command] [options]")

	fmt.Println("\nBASIC COMMANDS:")
	fmt.Println(" --help -h: Shows help menu")
	fmt.Println(" --create-service, -cS <destination_path>: Creates new Go-Web service in selected path")

	fmt.Println("\nPROJECT COMMANDS - RUNS THIS COMMANDS ONLY IN PROJECT ROOT!")
	fmt.Println(" --make-controller, -mC <controller_name>: Creates new Go-Web controller")
	fmt.Println(" --make-command, -mCMD <command_name>: Creates new Go-Web command")
	fmt.Println(" --make-model, -mM <model_name>: Creates new Go-Web model")
	fmt.Println(" --make-migration , -mDBM <migration_name>: Creates new Go-Web SQL migration")
	fmt.Println(" --make-middleware , -mMW <middleware_name>: Creates new middleware")
	fmt.Println(" --make-job, -mJ <job_name>: Creates new async job")
	fmt.Println(" --show-routes, -sR: Shows service routes")
	fmt.Println(" --migrate-up, -mU: Executes migrations")
	fmt.Println(" --migrate-rollback, -mR <steps>: Executes migrations rollback")
	fmt.Println(" --generate-key, -gK: Generate new application key")
	fmt.Println(" --http-load, -hL <target.json>: Executes HTTP load test based on json configuration")

	fmt.Println("\nPROJECT COMMANDS:")
	fmt.Println("WARNING: This commands are available in your project executable.")
	fmt.Println("To run the followings command you should compile your project and run:")
	fmt.Println("Usage: <your-binary> <command>:<option>")

	fmt.Println("\n database:seed <model_name>: Executes seeder (all available models if is not specified).")
	fmt.Println(" show:commands : Shows all custom cli commands")
	fmt.Println(" server:daemon : Run your project as a daemon")
	fmt.Println(" server:run : Run your project ")
}
