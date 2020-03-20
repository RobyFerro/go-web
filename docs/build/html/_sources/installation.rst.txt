Installation
############
You can install Go-Web as follows :

.. code-block:: bash

    mkdir <project_folder>
    go mod init <your_project_name>
    go get github.com/RobyFerro/go-web@v0.1.0-alpha
    cp $GOPATH/bin/go-web /usr/bin/goweb
    sudo go-web install <project_directory>
    sudo chown -R user:group <project_directory>
    sudo chmod -R +w <project_directory>

If you’re using Visual Studio Code you can install Go-Web framework with it’s specific extension:
https://marketplace.visualstudio.com/items?itemName=Roberto.go-web
