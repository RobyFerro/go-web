package middleware

import (
	"net/http"
)

// Middleware struct is extended by every middleware.
// To have and example see the below code.
type Middleware struct {
	Handler http.Handler
}
