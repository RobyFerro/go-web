CLI Interface
#############
Go-Web has a built-in command-line interface that makes easy for developers to use the framework.
Before using the CLI, the developer needs to compile the project by running command:
*go build goweb.go*

After compiling the project, the CLI can be used to view all commands supported by Go-Web
*./goweb show:commands*

The following listing table shows commands presented to the user by show:command:

.. image:: commands.png

Create custom commands
----------------------
Go-Web command line interface (CLI) can be extended by running command

*./goweb cmd:create <name>*

Before being available to Go-Web, commands must be registered in the console kernel at

*〈go-web〉/app/console/kernel.go*

The following listing shows the registration of command Greetings:

.. code-block:: go

    var (
        Commands = map[string]interface{} {
            /* Default commands */
            "migration:up": &command.MigrationUp{},
            "migration:create": &command.MigrationCreate{},
            "job:create": &command.JobCreate{},
            "install": &command.Install{},
            /* Default commands */
            /* Registration of Greetings command */
            "greetings": &command.Greetings{},
        }
    )

The command registration Commands variable is used by Go-Web to recognize and list supported commands.