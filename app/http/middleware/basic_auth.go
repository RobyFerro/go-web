package middleware

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

// BasicAuth used to check if the basic auth session is present.
// Use this middleware to protect resources with the basic authentication.
func (Middleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := kernel.Container.Invoke(func(s *sessions.CookieStore) {

			session, err := s.Get(r, "basic-auth")
			if err != nil {
				log.Fatal(err)
			}

			if session.Values["user"] == nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				next.ServeHTTP(w, r)
			}

		}); err != nil {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r)
	})
}
