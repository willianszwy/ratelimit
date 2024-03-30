package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
	"willianszwy/ratelimit/configs"
	"willianszwy/ratelimit/internal/ratelimit"
)

func main() {
	config, err := configs.LoadConfig("")
	if err != nil {
		panic(err)
	}
	log.Println("Load config...")
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword,
		DB:       0,
	})

	rlConfig := ratelimit.Config{
		IPMaxRequests:    int64(config.IPMaxRequests),
		IPBlockedTime:    time.Millisecond * time.Duration(config.IPBlockedTime),
		TokenMaxRequests: int64(config.TokenMaxRequests),
		TokenBlockedTime: time.Millisecond * time.Duration(config.TokenBlockedTime),
		TokenHeader:      config.TokenName,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(ratelimit.Middleware(rlConfig, rdb))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", r)
}
