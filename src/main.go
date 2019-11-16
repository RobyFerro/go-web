package main

import (
	"smartcharry/src/router"
	"smartcharry/src/service"
)

func main() {
	routes := router.WebRouter()
	service.StartHttpServer(routes)
}
