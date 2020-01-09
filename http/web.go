package http

import (
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
		hasMiddleware := route.Middleware != ""

		if hasMiddleware {
			subRouter := router.Path(route.Path).Subrouter()
			subRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
				in := []reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)}
				method.Call(in)
			}).Methods(route.Method)

			subRouter.Use(func(handler http.Handler) http.Handler {
				middleware := Middleware{Handler: handler}
				m := reflect.ValueOf(middleware)
				method := m.MethodByName(route.Middleware)
				test := method.Call([]reflect.Value{reflect.ValueOf(handler)})

				return test[0].Interface().(http.Handler)
			})
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				in := []reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)}
				method.Call(in)
			}).Methods(route.Method)
		}

	}
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

		for _, mw := range group.Middleware {
			subRouter.Use(func(handler http.Handler) http.Handler {
				middleware := Middleware{Handler: handler}
				m := reflect.ValueOf(middleware)
				method := m.MethodByName(mw)
				test := method.Call([]reflect.Value{reflect.ValueOf(handler)})

				return test[0].Interface().(http.Handler)
			})
		}
	}
}
