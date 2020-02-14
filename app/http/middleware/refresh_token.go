package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	"net/http"
)

// Refresh JWT token
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := gwf.ServiceContainer.Invoke(func(a *gwf.Auth) {
			a.RefreshToken()
		})

		if err != nil {
			gwf.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}
