package middleware

import "net/http"

// Middleware struct is extended by every middleware.
type Middleware struct {
	Handler http.Handler
}
