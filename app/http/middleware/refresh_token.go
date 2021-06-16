package middleware

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	gwf "github.com/RobyFerro/go-web/app/auth"
	"net/http"
)

// RefreshToken return a new token in request response
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var a gwf.JWTAuth
		kernel.Container.Invoke(func(conf *kernel.Conf) {
			a.RefreshToken(w, conf.App.Key)
		})

		next.ServeHTTP(w, r)
	})
}
