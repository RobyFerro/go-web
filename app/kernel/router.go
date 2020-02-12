package kernel

import (
	"fmt"
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/http/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

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
	var wg sync.WaitGroup
	wg.Add(len(mwList))

	for _, mw := range mwList {
		go func(mwr string, waitGroup *sync.WaitGroup) {
			m := reflect.ValueOf(middleware.Middleware{})
			method := m.MethodByName(mwr)

			callable := method.Interface().(func(handler http.Handler) http.Handler)
			midFunc = append(midFunc, callable)
			waitGroup.Done()
		}(mw, &wg)
	}
	wg.Wait()

	return midFunc
}

// Give access to public folder. With the /public prefix you can access to all of your assets.
// This is mandatory to access to public files (.js, .css, images, etc...).
func giveAccessToPublicFolder(router *mux.Router) {
	publicDirectory := http.Dir(config.GetDynamicPath("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(publicDirectory)))
}
