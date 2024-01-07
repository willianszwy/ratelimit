package middleware

import (
	"github.com/redis/go-redis/v9"
	"net/http"
)

// RateLimitMiddleware HTTP middleware setting a rate-limit to request
func RateLimitMiddleware(rdb *redis.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			next.ServeHTTP(w, r)
		})
	}
}
