package http

import (
	. "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/time/rate"
	"ikdev/go-web/app/config"
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
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := ServiceContainer.Invoke(func(a *helper.Auth) {
			a.RefreshToken()
		})

		if err != nil {
			exception.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}

// Check user authentication
func (Middleware) Auth(next http.Handler) http.Handler {
	var key string

	err := ServiceContainer.Invoke(func(c config.Conf) {
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

// Set a limit of request allowed in a specific time
func (Middleware) RateLimiter(next http.Handler) http.Handler {
	var limiter = rate.NewLimiter(1, 3)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
