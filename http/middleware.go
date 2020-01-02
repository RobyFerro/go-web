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

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	var key string

	err := Container.Invoke(func(c config.Conf) {
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

	return jwtMiddleware.Handler(next)
}

func RefreshTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := Container.Invoke(func(a *helper.Auth) {
			a.RefreshToken()
		})

		if err != nil {
			exception.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}
