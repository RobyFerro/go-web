# Go-Web
A simple lightweight web framework written in Go for Golang lovers!.
The project are currently in work in progress. Feel free to contribute!

## Configuration
Copy config.yml.example to new file named config.yml. You can customize your app by editing this file.

**WARNING**
To ensure that your authentication/authorization flow is completely secure you've to change "app.key" variables.

## Getting started
### Routing
Routes in Go-Web are handled by routing.yml file in root directory. This is an abstraction of [Gorilla Mux](https://github.com/gorilla/mux).
Every route are composed with:

* Path: URI of this route
* Action: Business logic (simple binder to your controller)
* Method: HTTP method (GET, POST, PUT, DELETE, PATCH)
* Middleware: Middleware for this route

You can regroup a set of routes by insert your route under "group" node

### Actions
Actions are stored in http package and its used to bind the HTTP router with a controller.
To create a new actions you've to extend the base Action structure as built-in actions (see "Main" action)

### Controllers
Controllers are the main responsible of the business logic.
You can find all controller into controller directory. Every controller must extends Controller structure.

### Authentication
Go-Web is dispatched with a built-in JWT Auth method. 
This integration checks credentials passed with an HTTP POST request with the contents of users table located in SQL database (see next chapter).

## Database
To customize your database connection you've to edit config.yml file in project root directory. This integration is an abstraction of [Gorm](https://gorm.io/)

### Models
Models are stored in database/model directory registered into database/models.go file.
This registration is mandatory to use Migration/Seed/Drop method by goweb CLI tool.

### Migrations
Every model should have a Migration method. This method it's just a wrapper of GORM auto migrate tool.

### Seeding
To create a seeder you've to create a Seed method inside your model. This method should insert something into the handled table.
(see the built-in models).

## Service container (Dependency Injection)
// Todo

## Other service
Go-Web is dispatched with a Redis and MongoDB integrations.

### Redis
// Todo

### MongoDB
// Todo

## CLI Tool
// Todo

## Commands
// Todo












