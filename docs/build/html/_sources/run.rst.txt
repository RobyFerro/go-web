Compile and run
===============

If you'd like to run Go-web in order to try your new implementation you can run the following command:

.. code-block:: bash.

    // Make sure to run this command in project root
    make gwf-run

To compile the entire project you just need to run:

.. code-block:: bash.

    // Make sure to run this command in project root
    make gwf-build

Then to start you can simply run of your Go-Web commands:

.. code-block:: bash.

    ./goweb server:run // Run server normally
    ./goweb server:daemon // Run server as a daemon

.. tip::
    Check *./goweb show:commands* to see all Go-Web available commands

The server will start listening on port defined in file config.yml.
