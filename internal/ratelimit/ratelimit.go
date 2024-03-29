package ratelimit

import (
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"net/http"
)

type RateLimit struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *RateLimit {
	return &RateLimit{rdb: rdb}
}

func (rl *RateLimit) Verify(r *http.Request) bool {
	log.Println(r.Method, r.Host, r.Header)
	return rand.Int()%2 == 0
}
