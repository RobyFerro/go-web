package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	. "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

// Auth checks if the JWT used by the request is valid.
// This middleware must be used only with JWT authentication and will not work with the basic auth.
func (Middleware) Auth(next http.Handler) http.Handler {
	var key string
	conf := foundation.RetrieveConfig()

	if len(conf.App.Key) == 0 {
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
