package middleware

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/app/auth"
	"net/http"
)

type RefreshTokenMiddleware struct {
	Name        string
	Description string
}

// Handle return a new token in request response
func (RefreshTokenMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var a auth.JWTAuth
		conf := foundation.RetrieveConfig()
		a.RefreshToken(w, conf.Key)

		next.ServeHTTP(w, r)
	})
}

// GetName returns the middleware name
func (m RefreshTokenMiddleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m RefreshTokenMiddleware) GetDescription() string {
	return m.Description
}

func NewRefreshTokenMiddleware() RefreshTokenMiddleware {
	return RefreshTokenMiddleware{
		Name:        "RefreshToken",
		Description: "Refresh JWT token",
	}
}
