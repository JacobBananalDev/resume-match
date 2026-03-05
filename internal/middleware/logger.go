package middleware

import (
	"log"
	"net/http"
	"time"
)

/*
responseWriter is a wrapper around http.ResponseWriter
that lets us capture the HTTP status code.
*/
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

/*
WriteHeader captures the status code when the handler writes it.
*/
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

/*
Logger is HTTP middleware that logs request details.

Example log output:

POST /analyze 200 15ms
GET /health 200 1ms
*/
func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// Wrap the ResponseWriter
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Call the next handler in the chain
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s %d %s",
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration,
		)
	})
}