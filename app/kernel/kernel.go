package kernel

import "C"
import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"ikdev/go-web/app/config"
	"ikdev/go-web/app/http/controller"
	"ikdev/go-web/app/http/middleware"
	"ikdev/go-web/exception"
	"ikdev/go-web/helper"
	"ikdev/go-web/service"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type HttpKernel struct {
	Models    []interface{}
	Container *dig.Container
}

// Parse routing structure and set every route
func WebRouter() *mux.Router {
	routes := config.ConfigurationWeb()
	router := mux.NewRouter()

	handleSingleRoute(routes.Routes, router)
	handleGroups(routes.Groups, router)
	giveAccessToPublicFolder(router)

	return router
}

// Start HTTP kernel
func StartKernel() *HttpKernel {
	Container = service.BuildContainer(Router)

	return &HttpKernel{
		Models:    Models,
		Container: Container,
	}
}

// Register Base Controller in every custom controller
func registerBaseControllers(res http.ResponseWriter, req *http.Request) []interface{} {
	if err := Container.Invoke(func(db *gorm.DB, c config.Conf, a *helper.Auth, s *sessions.CookieStore) {
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

	// Insert implementation
	// Es: Redis, Elastic, Mongo connections
	checkControllerIntegrations(&BaseController)

	// Register base controller in all controllers
	var controllers []interface{}
	for _, ctr := range Controllers {
		c := reflect.ValueOf(ctr).Elem().FieldByName("BaseController")
		c.Set(reflect.ValueOf(BaseController))
	}

	return controllers
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

// Handle single path parsing.
// This method it's used to parse every single path. If middleware is present a sub-router with will be created
func handleSingleRoute(routes map[string]config.Route, router *mux.Router) {
	var wg sync.WaitGroup
	wg.Add(len(routes))

	for _, route := range routes {
		go func(r config.Route) {
			hasMiddleware := len(r.Middleware) > 0
			directive := strings.Split(r.Action, "@")
			if hasMiddleware {
				subRouter := mux.NewRouter()
				subRouter.HandleFunc(r.Path, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					method.Call([]reflect.Value{})
				}).Methods(r.Method)

				subRouter.Use(parseMiddleware(r.Middleware)...)
				router.Handle(r.Path, subRouter)
			} else {
				router.HandleFunc(r.Path, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					method.Call([]reflect.Value{})
				}).Methods(r.Method)
			}

			wg.Done()
		}(route)
	}

	wg.Wait()
}

// Parse route groups.
func handleGroups(groups map[string]config.Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()
		var wg sync.WaitGroup
		wg.Add(len(group.Routes))

		for _, route := range group.Routes {
			go func(r config.Route) {
				directive := strings.Split(r.Action, "@")
				if len(r.Middleware) > 0 {
					nestedRouter := mux.NewRouter()
					fullPath := fmt.Sprintf("%s%s", group.Prefix, r.Path)
					nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
						cc := GetControllerInterface(directive, writer, request)
						method := reflect.ValueOf(cc).MethodByName(directive[1])
						method.Call([]reflect.Value{})
					}).Methods(r.Method)

					nestedRouter.Use(parseMiddleware(r.Middleware)...)
					subRouter.Handle(r.Path, nestedRouter)
				} else {
					subRouter.HandleFunc(r.Path, func(writer http.ResponseWriter, request *http.Request) {
						cc := GetControllerInterface(directive, writer, request)
						method := reflect.ValueOf(cc).MethodByName(directive[1])
						method.Call([]reflect.Value{})
					}).Methods(r.Method)
				}

				wg.Done()
			}(route)
		}

		wg.Wait()

		subRouter.Use(parseMiddleware(group.Middleware)...)
	}
}

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []string) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc
	for _, mw := range mwList {
		m := reflect.ValueOf(middleware.Middleware{})
		method := m.MethodByName(mw)

		callable := method.Interface().(func(handler http.Handler) http.Handler)
		midFunc = append(midFunc, callable)
	}

	return midFunc
}

// Get specific controller interface by its directive
func GetControllerInterface(directive []string, w http.ResponseWriter, r *http.Request) interface{} {
	var result interface{}

	// Insert request and response to every controller
	registerBaseControllers(w, r)

	// Find right controller
	for _, contr := range Controllers {
		controllerName := reflect.Indirect(reflect.ValueOf(contr)).Type().Name()
		if controllerName == directive[0] {
			result = contr
		}
	}

	return result
}

// Give access to public folder
// With the /public prefix you can access to all of your assets
func giveAccessToPublicFolder(router *mux.Router) {
	publicDirectory := http.Dir(config.GetDynamicPath("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(publicDirectory)))
}
