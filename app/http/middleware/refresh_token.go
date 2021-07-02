package middleware

import (
	foundation "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web/app/auth"
	"net/http"
)

// RefreshToken return a new token in request response
func (Middleware) RefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var a auth.JWTAuth
		err := foundation.RetrieveSingletonContainer().Invoke(func(conf *kernel.Conf) {
			a.RefreshToken(w, conf.App.Key)
		})

		if err != nil {
			return
		}

		next.ServeHTTP(w, r)
	})
}
