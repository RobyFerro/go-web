Controllers
===========
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

.. warning::
    After manually creating a controller, developers may need to manually register it: to do so, the controller needs to be added to Controllers list defined in
    *<go-web>/register.go*