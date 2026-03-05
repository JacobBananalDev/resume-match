package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

/*
RateLimiter limits how many requests can be processed
within a time window.

Example:
5 requests per second
*/
func RateLimiter(rps int) func(http.Handler) http.Handler {

	// Create a limiter allowing rps requests per second
	limiter := rate.NewLimiter(rate.Limit(rps), rps)

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Check if request is allowed
			if !limiter.Allow() {

				http.Error(
					w,
					"Too Many Requests",
					http.StatusTooManyRequests,
				)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}