Routing
============================
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