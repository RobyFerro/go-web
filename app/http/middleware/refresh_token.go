package middleware

import (
	"net/http"

	gwf "github.com/RobyFerro/go-web-framework"
)

// RefreshToken return a new token in request response
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var a gwf.Auth
		gwf.Container.Invoke(func(conf *gwf.Conf) {
			a.RefreshToken(w, conf.App.Key)
		})

		next.ServeHTTP(w, r)
	})
}
