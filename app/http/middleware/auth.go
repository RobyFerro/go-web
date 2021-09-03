package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	. "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	Name        string
	Description string
}

// Handle checks if the JWT used by the request is valid.
// This middleware must be used only with JWT authentication and will not work with the basic auth.
func (AuthMiddleware) Handle(next http.Handler) http.Handler {
	conf := foundation.RetrieveConfig()

	if len(conf.Key) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := New(Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtMiddleware.Handler(next)
}

// GetName returns the middleware name
func (m AuthMiddleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m AuthMiddleware) GetDescription() string {
	return m.Description
}

func NewAuthMiddleware() AuthMiddleware {
	mw := AuthMiddleware{
		Name:        "Auth",
		Description: "Provides JWT authentication",
	}

	return mw
}
