package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"ikdev/go-web/config"
	"net/http"
	"reflect"
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
		r := reflect.ValueOf(Action{})
		method := r.MethodByName(route.Action)
		hasMiddleware := len(route.Middleware) > 0

		if hasMiddleware {
			subRouter := mux.NewRouter()
			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				in := []reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)}
				method.Call(in)
			}).Methods(route.Method)

			var midFunc []mux.MiddlewareFunc
			for _, mw := range route.Middleware {
				m := reflect.ValueOf(Middleware{})
				method := m.MethodByName(mw)

				callable := method.Interface().(func(handler http.Handler) http.Handler)
				midFunc = append(midFunc, callable)
			}

			subRouter.Use(midFunc...)
			router.Handle(route.Path, subRouter)
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				in := []reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)}
				method.Call(in)
			}).Methods(route.Method)
		}
	}
	fmt.Println("Test")
}

// Parse route groups.
func handleGroups(groups map[string]config.Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()
		for _, route := range group.Routes {
			r := reflect.ValueOf(Action{})
			method := r.MethodByName(route.Action)

			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				in := []reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)}
				method.Call(in)
			}).Methods(route.Method)
		}

		var midFunc []mux.MiddlewareFunc
		for _, mw := range group.Middleware {
			m := reflect.ValueOf(Middleware{})
			method := m.MethodByName(mw)

			callable := method.Interface().(func(handler http.Handler) http.Handler)
			midFunc = append(midFunc, callable)
		}

		subRouter.Use(midFunc...)
	}
}
