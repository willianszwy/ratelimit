package ratelimit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	IPMaxRequests    int64
	IPBlockedTime    time.Duration
	TokenMaxRequests int64
	TokenBlockedTime time.Duration
	TokenHeader      string
}

type RateLimit struct {
	rdb    *redis.Client
	config Config
}

func New(config Config, rdb *redis.Client) *RateLimit {
	return &RateLimit{rdb: rdb, config: config}
}

func (rl *RateLimit) Verify(r *http.Request) bool {

	token := r.Header.Get(rl.config.TokenHeader)
	if token != "" {
		return rl.limitByToken(r.Context(), token)
	}
	return rl.limitByIp(r)
}

func (rl *RateLimit) limitByToken(ctx context.Context, token string) bool {
	log.Println("limiting by token")
	log.Println(rl.rdb.TTL(ctx, token).String())
	current, err := rl.rdb.Get(ctx, token).Int64()
	if err != redis.Nil && current > rl.config.TokenMaxRequests {
		log.Println("err", err, "count", current)
		return true
	}
	current, err = inc(ctx, rl.rdb, token, rl.config.TokenBlockedTime.Milliseconds())
	if err != nil {
		log.Println(err)
	}
	return false
}

func (rl *RateLimit) limitByIp(r *http.Request) bool {
	log.Println("limiting by ip")
	ctx := r.Context()
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	log.Println(rl.rdb.TTL(ctx, ip).String())
	current, err := rl.rdb.Get(ctx, ip).Int64()
	if err != redis.Nil && current > rl.config.IPMaxRequests {
		log.Println("err", err, "count", current)
		return true
	}
	current, err = inc(ctx, rl.rdb, ip, rl.config.IPBlockedTime.Milliseconds())
	if err != nil {
		log.Println(err)
	}
	return false
}

func inc(ctx context.Context, rdb *redis.Client, key string, expire int64) (int64, error) {
	incrBy := redis.NewScript(`
       local current
       local time = ARGV[1]
       current = redis.call("incr",KEYS[1])
       if current == 1 then
          redis.call("pexpire",KEYS[1],time)
       end
       return current
    `)

	keys := []string{key}
	values := []interface{}{expire}
	count, err := incrBy.Run(ctx, rdb, keys, values...).Int64()
	if err != nil {
		return 0, err
	}
	return count, nil
}
