package middleware

import (
	http2 "github.com/RobyFerro/go-web/app/http"
	"github.com/RobyFerro/go-web/exception"
	"github.com/RobyFerro/go-web/helper"
	"net/http"
)

// Refresh JWT token
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := http2.ServiceContainer.Invoke(func(a *helper.Auth) {
			a.RefreshToken()
		})

		if err != nil {
			exception.ProcessError(err)
		}

		next.ServeHTTP(w, r)
	})
}
