package main

import (
	"ikdev/smartcherry/router"
	"ikdev/smartcherry/service"
)

func main() {
	routes := router.WebRouter()
	service.StartHttpServer(routes)
}
