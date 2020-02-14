package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	. "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

// Check if the JWT used by the request is valid.
// This middleware must be used only with JWT authentication and will not work with the basic auth.
func (Middleware) Auth(next http.Handler) http.Handler {
	var key string

	err := gwf.ServiceContainer.Invoke(func(c *gwf.Conf) {
		key = c.App.Key
	})

	if err != nil {
		gwf.ProcessError(err)
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
