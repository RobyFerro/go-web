CLI Interface
#############
Go-Web has a built-in command-line interface that makes easy for developers to use the framework.
Before using the CLI, the developer needs to compile the project by running command:
*go build goweb.go*

After compiling the project, the CLI can be used to view all commands supported by Go-Web
*./goweb show:commands*

The following listing table shows commands presented to the user by show:command:

.. list-table:: Alfred available commands
    :widths: 50 50
    :header-rows: 1

    * - Commands
      - Descriptions
    * - database:seed <model_name>
      - Executes seeder (all available models if is not specified)
    * - show:commands
      - Shows all custom CLI commands
    * - server:daemon
      - Run Go-Web server as a daemon
    * - server:run
      - Run Go-Web server normally

Create custom commands
----------------------
Go-Web command line interface (CLI) can be extended by running command

.. code-block:: go

    alfred -mCMD <command_name>

Before being available to Go-Web, commands must be registered in *register.go*.
The following listing shows the registration of command Greetings:

.. warning::

    Command must be registered in the *register.go* file located in project root directory.

.. code-block:: go

	Commands = gwf.CommandRegister{
		List: map[string]interface{}{
			"queue:failed": &console.QueueFailed{},
			"queue:run":    &console.QueueRun{},
                        "greetings":    &console.Greetings{}, // new
			// Here is where you've to register your custom commands
		},
	}

The command registration Commands variable is used by Go-Web to recognize and list supported commands.