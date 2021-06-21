Configuration
#############
Configuring a service is straightforward and developers can use config.yml.example as an example for further customization; Go-Web will look for file config.yml, thus the aforementioned
config.yml.example can be copied with such name.
The following listing demonstrates how a developer can configure both MySQL database and HTTP server.

.. code-block:: yaml

    database:
      driver: "mysql"
      host: "localhost"
      port: 3306
      database: "your_database_name"
      user: "<db_username>"
      password: "<db_password>"
    server:
      name: "localhost"
      port: 8080
