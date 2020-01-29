<p align="center">
<img src="logo.png" alt="Go-Web">
</p>

## About Go-Web
A simple lightweight web framework written in Go for Golang lovers!.
The project are currently in work in progress. Feel free to contribute!

## Configuration
Copy config.yml.example to new file named config.yml. You can customize your app by editing this file.
If you need to add some extra configuration you've to update "Conf" structure in "configuration" package.

###### 1. Database
In this section you've to set your database connection

###### 2. Redis 
Set Redis connection. Mandatory for async jobs.

###### 3. MongoDB
Set MongoDB connection.

###### 4. Elastic Search
Set ElasticSearch connections

###### 5. Server
This section handle the main HTTP server configuration.

###### 6. App
Actually this section is used only to handle JWT key but you can use it to handle main app configuration.

###### 7. Mail (WIP)
Used to handle SMTP connection.

###### 8. Exception
Here you can find all exception implementation.

## Basic CLI commands
Go-Web is bundled with a numbers of helpful commands:

* `./goweb server:run` Run the HTTP server
* `./goweb server:daemon` Run the HTTP server as a daemon
* `./goweb migration:create <migration_name>` Create migrations .sql files
* `./goweb migration:up` Run the database migration
* `./goweb migration:rollback <steps>` Run the database migration
* `./goweb seed` Run the database seeder
* `./goweb queue:run <queue_name>` Run the selected job queue
* `./goweb queue:failed` Retry all failed jobs

You can implement your custom commands by adding your code into "command" package and register it into ./goweb.go switch statement.

## Routing
Routes in Go-Web are handled by routing.yml file in root directory. This is an abstraction of [Gorilla Mux](https://github.com/gorilla/mux).
Every route are composed with:

* Path: URI of this route
* Action: Is the coords of your business logic. The syntax is "ControllerName@MethodName".
* Method: HTTP method (GET, POST, PUT, DELETE, PATCH)
* Middleware: Middleware for this route

You can regroup a set of routes by insert your route under "group" node.

### Controllers
Controllers are the main responsible of the business logic.You can find all controller into "controller" package. 
Every controller must extends BaseController structure (which provides access to the service container) ad it must be registered in "register" method present into "http" package.

```
// This method will return an array of controllers.
// Used by Go-Web routing
// Every time you add a new controller you should register it in this method
func register(bc controller.BaseController) []interface{} {
	return []interface{}{
		&controller.UserController{BaseController: bc},
		&controller.AuthController{BaseController: bc},
		&controller.HomeController{BaseController: bc},
        // Here is where you've to register your custom controller
	}
}
```

### Structure of BaseController
The following structure is how BaseController is build.

````
type BaseController struct {
	DB       *gorm.DB               // Provide access to MySql instance
	Response http.ResponseWriter    // HTTP response
	Request  *http.Request          // HTTP request
	Config   config.Conf            // Go-Web configuration
	Auth     *helper.Auth           // Authentication/Authorization method
	Redis    *redis.Client          // Provide access to Redis instance
	Mongo    *mongo.Database        // Provide access to MongoDB instance
	Elastic  *elasticsearch.Client  // Provide access to ElasticSearch instance
}
````

Extending the BaseController every controller can access to all resources by a simple call. Es:

````
//Used to check authenticated user
func (c *UserController) Profile() {
	c.Auth.GetUser(c.Request)

	_, _ = c.Response.Write([]byte("Authorization ok"))
}
````

As you can see we've got access to Auth methods simply calling "c.Auth".

### Authentication
Go-Web is dispatched with a built-in JWT Auth method. 
To get a new token you must follow this flow:

1. Execute user login by sending credentials to "/login" route with POST method. 
2. If previous request has succeeded a new JWT Token will be return [[1]](https://jwt.io/introduction/).
3. Include the given token in the future requests header. Es: Authentication: Bearer <your JWT token>

This authentication requires that you've used the basic database boilerplate bundled with "Go-Web" (see next section).

## Database
First of all you've to set database connection inside "config.yml" file (at the time only MySql is supported).
This integration is an abstraction of [Gorm](https://gorm.io/).

Database instance is available by default inside your controller (by implementing a service container).

### Models
Models are stored in "model" package and registered into "database/models.go" file.
If you need to add a new model this bust be registered into "GetModels" method (database package). 

````
// This method will return an array of models.
// Used to handle migration, seeding and drop operations.
// Every time you add a new model you should register it in this method
func GetModels() []interface{} {
	return []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}
}
````

Every model must extends "gorm.Model" and implements "Migrate, Drop and Seed" method.

#### Migration
Migration are now handled by simple .sql file. You can create new migration by `./goweb migration:create <migration_name>`.
This command will generate two kind of file: ".up.sql" and ".down.sql."

* The ".up.sql" file will be used to migrate.
* The ".down.sql" file will be used to rollback.

#### Seeding
By implementing "Seed" method you'te able to seed your table with a "fake" data.
Go-Web uses [gofakeit](https://github.com/brianvoe/gofakeit) as faker library.

See the below example:
````
// Execute model seeding
func (User) Seed(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		password := gofakeit.Password(true, true, true, true, false, 32)
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

		user := User{
			Name:     gofakeit.FirstName(),
			Surname:  gofakeit.LastName(),
			Username: gofakeit.Username(),
			Password: string(encryptedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			exception.ProcessError(err)
		}
	}
}
````

## Async jobs
Jobs in Go-Web are handled by a simple Redis List. To register a new Job you must extend Job structure in job directory. 
You can use Schedule and Execute method to handle it. 

**WARNING** To run schedule jobs you must start queue (see CLI commands) 

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

## Error handling
Every error in Go-Web should be handled by "ProcessError" in "exception" package and stored in "storage/log/error.log" file. 
If you've implemented Sentry URI in "config.yml" all error will be sent to your Sentry account.

## Docker integration
A Docker file and docker-compose.yml are already provided in Go-Web. Feel free to edit docker-compose.yml to customize your services. 
To simplify docker configuration the "env.example" file contains basic service configuration. Copy and edit "env.example" content to a new ".env" file.

## Custom commands
You can create custom command by creating new structure inside "command" package.
Every struct must have a "Signature string" property and a "Run" method. 
This method must have the following signature: 

````
func (c *ServerDaemon) Run(sc *dig.Container, args string) {
    // Your logic
}
````

Thanks to "sc" parameter you have complete access to your service container and with "args" to the argument passed in your command.
Last but not least you've to register the new created command in "GetCommands" method (kernel.go).

````
// This method will return an array of commands.
// Used by Go-Web routing
// Every time you add a new controller you should register it in this method
func GetCommands() map[string]interface{} {
	return map[string]interface{}{
		"migration:up":       &MigrationUp{Signature: "migration:up"},
		"migration:create":   &MigrationCreate{Signature: "migration:create"},
		"migration:rollback": &MigrateRollback{Signature: "migration:rollback"},
		"queue:failed":       &QueueFailed{Signature: "queue:failed"},
		"queue:run":          &QueueRun{Signature: "queue:run"},
		"seed":               &Seeder{Signature: "seed"},
		"server:daemon":      &ServerDaemon{Signature: "server:daemon"},
		"server:run":         &ServerRun{Signature: "server:run"},
		// Here is where you've to register your custom command
	}
}
````

