HTTP load test
______________

Go-Wep provides a simple way to test our HTTP server.

1. Setup a JSON configuration file
----------------------------------
First of all we need do configure a json-file as follow:

.. code-block:: json

    {
        "routes": [
            {
                "method": "POST",
                "url": "/users",
                "Body": "request.test.json",
                "Header": "application/json",
                "Rate": 100,
                "Time": 2
            },
            {
                "method": "GET",
                "url": "",
                "Body": "",
                "Header": "application/json",
                "Rate": 100,
                "Time": 2
            }
        ]
    }

Of course you can customize every routes.

2. Run the test
---------------
Once created the configuration file we can start our stress test by typing:

.. code-block:: bash

    alfred -hL <configuration.json>





