package middleware

import (
	. "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"ikdev/go-web/app/config"
	http2 "ikdev/go-web/app/http"
	"ikdev/go-web/exception"
	"log"
	"net/http"
)

// Check user authentication
func (Middleware) Auth(next http.Handler) http.Handler {
	var key string

	err := http2.ServiceContainer.Invoke(func(c config.Conf) {
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
