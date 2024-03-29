package main

import (
	"github.com/redis/go-redis/v9"
	"net/http"
	middleware2 "willianszwy/ratelimit/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware2.RateLimitMiddleware(rdb))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", r)
}
