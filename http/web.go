package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"ikdev/go-web/config"
	"net/http"
	"reflect"
	"strings"
)

// Parse routing structure and set every route
func WebRouter() *mux.Router {
	routes := config.ConfigurationWeb()
	router := mux.NewRouter()

	handleSingleRoute(routes.Routes, router)
	handleGroups(routes.Groups, router)

	return router
}

// Handle single path parting.
// This method it's used to parse every single path. If middleware is present a subrouter with will be created
func handleSingleRoute(routes map[string]config.Route, router *mux.Router) {
	for _, route := range routes {
		hasMiddleware := len(route.Middleware) > 0
		directive := strings.Split(route.Action, "@")
		if hasMiddleware {
			subRouter := mux.NewRouter()
			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				cc := GetControllerInterface(directive, writer, request)
				method := reflect.ValueOf(cc).MethodByName(directive[1])
				method.Call([]reflect.Value{})
			}).Methods(route.Method)

			subRouter.Use(parseMiddleware(route.Middleware)...)
			router.Handle(route.Path, subRouter)
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				cc := GetControllerInterface(directive, writer, request)
				method := reflect.ValueOf(cc).MethodByName(directive[1])
				method.Call([]reflect.Value{})
			}).Methods(route.Method)
		}
	}
}

// Parse route groups.
func handleGroups(groups map[string]config.Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()
		for _, route := range group.Routes {
			directive := strings.Split(route.Action, "@")
			if len(route.Middleware) > 0 {
				nestedRouter := mux.NewRouter()
				fullPath := fmt.Sprintf("%s%s", group.Prefix, route.Path)
				nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					method.Call([]reflect.Value{})
				}).Methods(route.Method)

				nestedRouter.Use(parseMiddleware(route.Middleware)...)
				subRouter.Handle(route.Path, nestedRouter)
			} else {
				subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					method.Call([]reflect.Value{})
				}).Methods(route.Method)
			}
		}

		subRouter.Use(parseMiddleware(group.Middleware)...)
	}
}

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []string) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc
	for _, mw := range mwList {
		m := reflect.ValueOf(Middleware{})
		method := m.MethodByName(mw)

		callable := method.Interface().(func(handler http.Handler) http.Handler)
		midFunc = append(midFunc, callable)
	}

	return midFunc
}

// Get specific controller interface by its directive
func GetControllerInterface(directive []string, w http.ResponseWriter, r *http.Request) interface{} {
	var controller interface{}
	for _, contr := range GetControllers(w, r) {
		controllerName := reflect.Indirect(reflect.ValueOf(contr)).Type().Name()
		if controllerName == directive[0] {
			controller = contr
		}
	}

	return controller
}
