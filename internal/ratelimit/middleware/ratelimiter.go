package ratelimit

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
	"willianszwy/ratelimit/internal/ratelimit"
)

var errTooManyRequests = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")

// Middleware RateLimit HTTP middleware setting a rate-limit to request
func Middleware(rdb *redis.Client) func(next http.Handler) http.Handler {
	rl := ratelimit.New(rdb)
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
