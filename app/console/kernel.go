package console

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/service"
)

var (
	// Commands is used to register all console commands.
	Commands = register.CommandRegister{}
	// Services will be used to create the Console Service Container.
	// This container is created to allow dependency injection through console commands.
	Services = register.ServiceRegister{
		service.ConnectDB,
		service.ConnectElastic,
		service.ConnectMongo,
		service.ConnectRedis,
	}
)
