Installation
############
Standard installation
---------------------
You can download and install Go-Web by following these steps:
    * Download Go-Web release from `GitHub <https://github.com/RobyFerro/go-web/releases>`__
    * Extract the content in you project root
    * Rename the `config.yaml.example` file in the project root by removing the `.example` suffix
    * Download all dependencies by executing `go mod download` in project root
    * Execute `go run . show:commands` to see all available GWF command

Docker
---------------------
Go-Web provides a docker-compose.yml file that allows developers to easily set up a new development environment: this requires both Docker and Docker-compose installed on the development system.

.. note::
    The docker-compose.yml defines several services, i.e. it is configured for providing instances ofMySQL, Redis, MongoDB and ElasticSearch; if needed, instances of other services may be added by modifying the docker-compose.yml file.




