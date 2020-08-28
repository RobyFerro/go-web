Run HTTP Server
===============
After creating a controller, registering it in the kernel and creating the binding with a route, the developer can run the server to ensure that the solution works; running a server can be done by running Go-Web like:

.. code-block:: bash.

    ./goweb server:run // Run server normally
    ./goweb server:daemon // Run server as a daemon

The server will start listening on port defined in file config.yml.
