<p align="center">
<img src="logo.png" alt="Go-Web">
</p>

A simple MVC web framework for Golang lovers!. 
## About Go-Web
Go-Web adopts a “convention over configuration” approach similarly to frameworks like Laravel ,Symfony and Rails.

By following this principle, Go-Web reduces the need for “repetitive” tasks like explicitly binding routes, actions and middleware; while not problematic per se, programmers may want to adopt “ready-to-use” solutions, for instances if they want easily build complex services, aiming at a reduced “productive lag”.

Programmers may want to use existing frameworks like Gin-Gonic and Go Buffalo, but Go-Web differs from them because of the aforementioned “convention over configuration” approach.

[See full documentation](https://goweb.readthedocs.io/en/latest/)

### Installation
Alfred is the command-line interface included with Go-Web. It provides a number of helpful commands that can assist you while you build your application.

#### Compiling from source
In order to compile "Alfred" from sources, make sure that:

* You have a correctly configure Go >= 1.13 environment
* $GOPATH is defined and $GOPATH/bin is in $PATH

Once you've met this conditions, you can run the following commands to compile and install Alfred in `/usr/local/bin/alfred`

```
go get github.com/RobyFerro/go-web
cd $GOPATH/src/github.com/RobyFerro/go-web
sudo make build
sudo make install
```

Then you can type `alfred -h` to see all available commands.

#### Create new service
To create a new Go-Web Service you've just to type:
`alfred -cS <selected_path>`




