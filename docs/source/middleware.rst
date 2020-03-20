Middleware
==========
Like controllers, a middleware can be created with command:

.. highlights:: ./goweb middleware:create <middleware name>

For instance, middleware named “Passthrough” can be created by running command:

.. highlights:: ./goweb middleware:create passthrough

After executing the command, the newly created middleware will be available in folder:

.. highlights:: <go-web>/app/http/middleware

As described previously, middlewares can be used for pre-processing requests received by routes defined in file routing.yml.