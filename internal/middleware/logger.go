package middleware

import (
	"log"
	"net/http"
	"time"
)

/*
Logger is HTTP middleware that logs request details.

Example log output:

POST /analyze 200 15ms
GET /health 200 1ms
*/
func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}