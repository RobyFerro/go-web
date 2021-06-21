HTTP
####
Go-Web encapsulate `Gorilla Mux <https://github.com/gorilla/mux>`_ to handles every HTTP requests. A `routing.yml` contains
the definitions of every route.


Routing
-------
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

Controllers
-----------
Being a MVC framework, Go-Web encourages the use of controllers, i.e. containers of the business logic of the application.
For instance, the controller named “SampleController” can be created by running command:

.. code-block:: bash

    alfred -mC sample_controller


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

Middleware
----------
Like controllers, a middleware can be created with command:

.. highlights:: ./goweb middleware:create <middleware name>

For instance, middleware named “Passthrough” can be created by running command:

.. highlights:: ./goweb middleware:create passthrough

After executing the command, the newly created middleware will be available in folder:

.. highlights:: <go-web>/app/http/middleware

| As described previously, middlewares can be used for pre/post processing requests.

.. note::
    Check `Gorilla Mux Middleware <https://github.com/gorilla/mux#middleware>`_ definition to more info about middlewares.

Authentication
--------------
By default, Go-Web provides two ways for authenticating users:

*  JWT-based authentication
*  basic (base) authentication

JWT Authentication
~~~~~~~~~~~~~~~~~~
Commonly used to authenticate users thought mobile applications or a SPA, JWT authentication is implemented by function JWTAuthentication of controller AuthController or, in Go-Web terms,
by endpoint **AuthController@JWTAuthentication**

The JSON structure used to represent credentials of a user must conform to JSON

.. code-block:: none

    {
        "username": <string, mandatory>,
        "password": <string, mandatory>
    }

The result of a successful login attempt with this type of authentication is a HTTP response containing a JWT token.
Resource access can be restricted only to authenticated users by adding middleware Auth to specific routes.

Basic authentication
~~~~~~~~~~~~~~~~~~~~~~~~~~~
Basic, or base, authentication is the simplest way to authenticate users for service access; this method is implemented by endpoint
**AuthController@BasicAuth**

The base authentication requires the same data structure as JWT-based method and routes can be protected by using middleware BasicAuth.


