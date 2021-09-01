package console

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/service"
)

var (
	// Commands is used to register all console commands.
	Commands = register.CommandRegister{
		List: map[string]interface{}{
			// Here is where you've to register your custom commands
		},
	}
	// Services will be used to create the Console Service Container.
	// This container is created to allow dependency injection through console commands.
	Services = register.ServiceRegister{
		List: []interface{}{
			service.ConnectDB,
			service.ConnectElastic,
			service.ConnectMongo,
			service.ConnectRedis,
		},
	}
)
