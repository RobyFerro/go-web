package http

import (
	. "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
	"ikdev/go-web/helper"
	"log"
	"net/http"
)

// Middleware struct is extended by every middleware.
// To have and example see the below code.
type Middleware struct {
	Handler http.Handler
}

// Log every actions printing used route
func (Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// Refresh JWT token
func (Middleware) RefreshToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := container.Invoke(func(a *helper.Auth) {
			a.RefreshToken()
		})

		if err != nil {
			exception.ProcessError(err)
		}

		Middleware{}.Handler.ServeHTTP(w, r)
	})
}

// Check user authentication
func (Middleware) Auth() http.Handler {
	var key string

	err := container.Invoke(func(c config.Conf) {
		key = c.App.Key
	})

	if err != nil {
		exception.ProcessError(err)
	}

	if len(key) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := New(Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtMiddleware.Handler(Middleware{}.Handler)
}
