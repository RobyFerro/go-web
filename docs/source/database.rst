Database
========

Models
------
In MVC frameworks models are responsible of the database interaction logic. Go-Web take advantage of GORM library to provide them.

To create a new model you can use its specific CLI command:

.. highlights:: ./goweb model:create <model name>

Models are located in: *<go-web>/database/model*

Migration
_________
Migrations are like version control for your database, allowing your team to easily modify and share the application’s database schema.
Developers can creates new migration as follows:

.. highlights:: ./goweb migration:create <migration_name>

Developer can find its newly created migration files in <go-web>/database/migration directory.


