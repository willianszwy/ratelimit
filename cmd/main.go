package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
	"net/http"
	ratelimit "willianszwy/ratelimit/internal/ratelimit/middleware"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(ratelimit.Middleware(rdb))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", r)
}
