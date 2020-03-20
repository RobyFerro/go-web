Controllers
===========
Being a MVC framework, Go-Web encourages the use of controllers, i.e. containers of the business logic of the application.
For instance, the controller named “SampleController” can be created by running command:

*./goweb controller:create sample*

Go-Web will create the the .go file containing controller named “SampleController” in folder:

*<go-web>/app/http/controller*

The content of the newly created file will be:

.. code-block:: go
    :caption: Sample controller

    package controller

    import "github.com/RobyFerro/go-web-framework"

    type SampleController struct{
        gwf.BaseController
    }

    // Main controller method
    func (c *SampleController) Main(){
        // Insert your custom logic
    }

When creating a controller, Go-Web will add to it the function Main, which could be expanded with some logic,
as shown in listing 4; controllers can be extended by adding new public functions.

.. code-block:: go
    :caption: List 4: Sample controller

    package controller

    import (
        "github.com/RobyFerro/go-web-framework"
        "github.com/RobyFerro/go-web/exception"
    )

    type SampleController struct{
        gwf.BaseController
    }

    // Main controller method
    func (c *SampleController) Main() {
        _, err := c.Response.Write([]byte("Hello world")) if err != nil {
            exception.ProcessError(err)
        }
    }

To gain access to everything underlying a Go-Web controller, including HTTP request and response, a controller needs to extend gwf.BaseController.
Because the service container is used to “resolve” all controllers in Go-Web, developers can type- hint any of their dependency because they will be injected into the controller instance, as represented by the following code:

.. code-block:: go
    :caption: List 5: Dependency injection in controller

    package controller

    import (
        "github.com/RobyFerro/go-web-framework" "github.com/RobyFerro/go-web/database/model" "github.com/jinzhu/gorm"
    )

    type SampleController struct{
        gwf.BaseController
    }

    // Main controller method
    func (c *SampleController) Main(db *gorm.DB) {
        var user model.User
        if err := db.Find(&user).Error;err != nil {
            gwf.ProcessError(err)
        }
    }


**Note**: both listings 4 and 5 includes a call to gwf.ProcessError(err); this is how Go-Web can handle errors, but developers may adopt another approach.

Registering a controller
------------------------
After creating a controller, developers need to register it into Go-Web kernel: to do so, the controller needs to be added to Controllers list defined in

*<go-web>/app/kernel/conf.go*

Binding controller to routes
----------------------------
Updating routes is simple and requires little changes to routing.yml file, which is located in the root folder of the project.
The definition of a route is, in fact, straightforward and routes can be organized in groups.

A route is defined by:

* path
    * describes the URI of the route
    * a path is expressed as a string which could define parameters and supports regular expressions as gorilla/mux does
    * requests targeting undefined routes will cause a “client error” response with HTTP status 404 (not found)
    * example: *‘‘/hello-world’’*
* action
    * describes the destination of a route as a combination of a controller and one of its functions
    * an action is expressed as the string <controller name>@<function name>
    * if the action cannot be resolved (undefined controller or action), Go-Web will produce an error
    * example: **SampleController@Main**
* method
    * describes the HTTP verb supported by the route
    * a method must be one of the verbs supported by HTTP, i.e.:
        * GET
        * HEAD
        * POST
        * PUT
        * DELETE
        * CONNECT
        * OPTIONS
        * TRACE
        * PATCH
    * requests targeting an existing route with a wrong method (i.e. one that is not supported by the route) will cause a “client error” response with HTTP status 405 (method not allowed)
* middleware (optionals)
    * represents the ordered list of middlewares that will process the request received by the route before performing the route’s action
    * the value of this property is a yml list of strings which must identify existing middlewares
    * example: Logging
* descriptions (optionals)
    * a string describing the purpose of the route
    * example: Returns JSON {‘‘message’’: ‘‘Hello World’’}