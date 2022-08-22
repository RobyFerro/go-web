package module

import (
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web/service"
)

var MainModule = register.DIModule{
	Name: "auth",
	Provides: []interface{}{
		service.ConnectDB,
		service.ConnectRedis,
		service.ConnectElastic,
		service.ConnectMongo,
	},
}
