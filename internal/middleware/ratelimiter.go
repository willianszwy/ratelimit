package middleware

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"willianszwy/ratelimit/ internal/ratelimit"
)

var errTooManyRequests = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")

// RateLimitMiddleware HTTP middleware setting a rate-limit to request
func RateLimitMiddleware(rdb *redis.Client) func(next http.Handler) http.Handler {
	rl := ratelimit.New(rdb)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("teste")
			if rl.Verify() {
				http.Error(w, errTooManyRequests.Error(), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
