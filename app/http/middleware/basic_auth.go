package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	"log"
	"net/http"
)

type BasicAuthMiddleware struct {
	Name        string
	Description string
}

// Handle used to check if the basic auth session is present.
// Use this middleware to protect resources with the basic authentication.
func (BasicAuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store := foundation.RetrieveCookieStore()
		session, err := store.Get(r, "basic-auth")
		if err != nil {
			log.Fatal(err)
		}

		if session.Values["user"] == nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}

		next.ServeHTTP(w, r)
	})
}

// GetName returns the middleware name
func (m BasicAuthMiddleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m BasicAuthMiddleware) GetDescription() string {
	return m.Description
}

func NewBasicAuthMiddleware() BasicAuthMiddleware {
	return BasicAuthMiddleware{
		Name:        "BasicAuth",
		Description: "Provides basic authentication",
	}
}
