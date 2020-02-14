package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/gorilla/sessions"
	"net/http"
)

// Used to check if the basic auth session is present.
// Use this middleware to protect resources with the basic authentication.
func (Middleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := gwf.Container.Invoke(func(s *sessions.CookieStore) {

			session, err := s.Get(r, "basic-auth")
			if err != nil {
				gwf.ProcessError(err)
			}

			if session.Values["user"] == nil {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte("Forbidden!"))
			}

		}); err != nil {
			gwf.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}
