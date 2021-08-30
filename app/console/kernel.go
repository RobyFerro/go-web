package console

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/app"
	"github.com/RobyFerro/go-web/service"
)

var (
	Commands = register.CommandRegister{
		List: map[string]interface{}{
			// Here is where you've to register your custom commands
		},
	}
	// Services will be used to create the Console Service Container.
	// This container is created to allow dependency injection through console commands.
	Services = register.ServiceRegister{
		List: []interface{}{
			app.Configuration,
			service.ConnectDB,
			service.ConnectElastic,
			service.ConnectMongo,
			service.ConnectRedis,
		},
	}
)
