# Go-Web
A simple lightweight web framework written in Go for Golang lovers!.
The project are currently in work in progress. Feel free to contribute!

## Configuration
Copy config.yml.example to new file named config.yml. You can customize your app by editing this file.

**WARNING**
To ensure that your authentication/authorization flow is completely secure you've to change "app.key" variables.

## Getting started
To run webserver `./goweb run:server` 

### Routing
Routes in Go-Web are handled by routing.yml file in root directory. This is an abstraction of [Gorilla Mux](https://github.com/gorilla/mux).
Every route are composed with:

* Path: URI of this route
* Action: Business logic (simple binder to your controller)
* Method: HTTP method (GET, POST, PUT, DELETE, PATCH)
* Middleware: Middleware for this route

You can regroup a set of routes by insert your route under "group" node

### Controllers
Controllers are the main responsible of the business logic.
You can find all controller into controller directory. Every controller must extends Controller structure.

### Authentication
Go-Web is dispatched with a built-in JWT Auth method. 
This integration checks credentials passed with an HTTP POST request with the contents of users table located in SQL database (see next chapter).

## Database
To customize your database connection you've to edit config.yml file in project root directory. This integration is an abstraction of [Gorm](https://gorm.io/)

### Models
Models are stored in database/model directory and registered into database/models.go file.
This registration is mandatory to use Migration/Seed/Drop method by goweb CLI tool.

### Migrations
Every model should have a Migration method. This method it's just a wrapper of GORM auto migrate tool.
To run migration `./goweb migrate`

### Seeding
To create a seeder you've to create a Seed method inside your model. This method should insert something into the handled table.
(see the built-in models).
To run seeder `./goweb seed`

## Schedule async jobs
Jobs in Go-Web are handled by a simple Redis List. 
To register a new Job you must extend Job structure in job directory. 
You can use Schedule and Execute method to handle it. 

**WARNING** To run schedule jobs you must start queue 

Example of jobs schedule:
```

// Schedule new job
data := job.MailStruct{
	To: []string{
		"test@test.com",
		"test@test1.com",
	},
	Message: "Hello world",
	}

serializedPayload, _ := json.Marshal(data)

j := job.Job{
	Name:       "Send email",
	MethodName: "Mail",
	Params:     job.Param{
		Name:    "message",
		Payload: string(serializedPayload),
		Type:    "int",
	},
}

j.Schedule("default", c.Redis)
// End schedule
```

## Service container (Dependency Injection)
Go-Web implement [Uber dig](https://github.com/uber-go/dig) library. This allow to use service container and dependency injection everything.
Service container is declared into service/container.go. Feel free to add your custom dependency.

## Other service
Go-Web is dispatched with:
* MongoDB
* Redis












