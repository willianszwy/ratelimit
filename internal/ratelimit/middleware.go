package ratelimit

import (
	"errors"
	"net/http"
)

var errTooManyRequests = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")

// Middleware RateLimit HTTP middleware setting a rate-limit to request
func Middleware(config Config, ks KeyStorage) func(next http.Handler) http.Handler {
	rl := New(config, ks)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if rl.Verify(r) {
				http.Error(w, errTooManyRequests.Error(), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
