package kernel

import "C"
import (
	"encoding/gob"
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/http/controller"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/exception"
	"github.com/RobyFerro/go-web/service"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"net/http"
	"reflect"
)

type HttpKernel struct {
	Models    []interface{}
	Container *dig.Container
}

// Return an HTTP kernel instance.
func StartKernel() *HttpKernel {
	Container = service.BuildContainer(Router)

	// Register user model struct to allow easy session encoding/decoding
	// Used only by the basic authentication.
	gob.Register(model.User{})

	return &HttpKernel{
		Models:    Models,
		Container: Container,
	}
}

// Returns a specific controller instance by comparing "directive" parameter with controller name.
func GetControllerInterface(directive []string, w http.ResponseWriter, r *http.Request) interface{} {
	var result interface{}

	// Insert request and response to every controller
	//registerBaseControllers(w, r)

	// Find the right controller
	for _, contr := range Controllers {
		controllerName := reflect.Indirect(reflect.ValueOf(contr)).Type().Name()
		if controllerName == directive[0] {
			registerBaseController(w, r, &contr)
			result = contr
		}
	}

	return result
}

// Parse a controller instance and implement it with the current base controller.
// This operation will give you access to all basic controller properties.
func registerBaseController(res http.ResponseWriter, req *http.Request, controller *interface{}) *interface{} {
	setBaseController(res, req)
	checkControllerIntegrations(&BaseController)

	c := reflect.ValueOf(*controller).Elem().FieldByName("BaseController")
	c.Set(reflect.ValueOf(BaseController))

	return controller
}

// Setting up the base controller instance (defined in conf.go) with the properties/method defined in the Service Container.
// Here you can implement the BaseController content.
// Remember to update even the structure (app/http/controller/controller.go)
func setBaseController(res http.ResponseWriter, req *http.Request) {
	if err := Container.Invoke(func(db *gorm.DB, c config.Conf, a *service.Auth, s *sessions.CookieStore) {
		BaseController = controller.BaseController{
			DB:       db,
			Response: res,
			Request:  req,
			Config:   c,
			Auth:     a,
			Session:  s,
		}
	}); err != nil {
		exception.ProcessError(err)
	}
}

// Check controller integrations
// Es: Redis, Elastic, Mongo connections
func checkControllerIntegrations(base *controller.BaseController) {

	// If is configured MongoDB will be implemented into service container
	if len(Configuration.Mongo.Host) > 0 {
		if err := Container.Invoke(func(m *mongo.Database) {
			base.Mongo = m
		}); err != nil {
			exception.ProcessError(err)
		}
	}

	// If is configured Redis will be implemented into service container
	if len(Configuration.Redis.Host) > 0 {
		if err := Container.Invoke(func(r *redis.Client) {
			base.Redis = r
		}); err != nil {
			exception.ProcessError(err)
		}
	}

	// If is configured ElasticSearch will be implemented into service container
	if len(Configuration.Elastic.Hosts) > 0 {
		if err := Container.Invoke(func(e *elasticsearch.Client) {
			base.Elastic = e
		}); err != nil {
			exception.ProcessError(err)
		}
	}
}
