package middleware

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct {
	Name        string
	Description string
}

// Handle log every actions printing used route
func (LoggingMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
