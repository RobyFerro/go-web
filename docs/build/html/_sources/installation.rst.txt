Installation
############
Alfred is the command-line interface included with Go-Web. It provides a number of helpful commands that can assist you while you build your application.

Compiling from source
---------------------
In order to compile "Alfred" from sources, make sure that:

* You have a correctly configure Go >= 1.13 environment
* $GOPATH is defined and $GOPATH/bin is in $PATH

Once you've met this conditions, you can run the following commands to compile and install Alfred in `/usr/local/bin/alfred`

.. code-block:: bash

    go get github.com/RobyFerro/go-web
    cd $GOPATH/src/github.com/RobyFerro/go-web
    sudo make build
    sudo make install

Then you can type `alfred -h` to see all available commands.

Create new service
------------------
To create a new Go-Web Service you've just to type:

.. code-block:: bash

    alfred -cS <selected_path>

If you’re using Visual Studio Code you can install Go-Web framework with it’s specific extension:
https://marketplace.visualstudio.com/items?itemName=Roberto.go-web
