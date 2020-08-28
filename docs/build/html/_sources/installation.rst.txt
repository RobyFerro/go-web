Installation
############
To install and create a new Go-Web Service you'll need to compile and use  :ref:`alfred-reference`.

Compiling from source
---------------------
In order to compile "Alfred" from sources, make sure that:

* You have a correctly configure Go >= 1.13 environment
* $GOPATH is defined and $GOPATH/bin is in $PATH

Linux and OSX
-------------
Once you've met this conditions, you can run the following commands to compile and install Alfred in `/usr/local/bin/alfred`

.. code-block:: bash

    go get github.com/RobyFerro/go-web
    cd $GOPATH/src/github.com/RobyFerro/go-web
    sudo make build
    sudo make install

Then you can type `alfred -h` to see all available commands.

Windows
-------
To install Go-Web on Windows system you'll need to manually compile Alfred and add the path to system env variables.

.. code-block:: bash

    go build $GOPATH/src/github/RobyFerro/go-web/tool/alfred.go

Create new service
------------------
To create a new Go-Web Service you've just to type:

.. code-block:: bash

    alfred -cS <selected_path>

.. note::
    If you’re using Visual Studio Code you can install Go-Web framework with it’s specific extension:
    https://marketplace.visualstudio.com/items?itemName=Roberto.go-web

    *Fully tested on Linux/OSX system*
