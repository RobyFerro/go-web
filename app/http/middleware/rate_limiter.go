package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
)

// Set a limit of request allowed in a specific time
func (Middleware) RateLimiter(next http.Handler) http.Handler {
	var limiter = rate.NewLimiter(1, 3)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
