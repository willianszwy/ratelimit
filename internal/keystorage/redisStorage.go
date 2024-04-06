package keystorage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisStorage struct {
	client *redis.Client
}

func New(rdb *redis.Client) *RedisStorage {
	return &RedisStorage{
		client: rdb,
	}
}

func (rd *RedisStorage) TTL(ctx context.Context, key string) string {
	return rd.client.TTL(ctx, key).String()
}

func (rd *RedisStorage) Get(ctx context.Context, key string) (int64, error) {
	r, err := rd.client.Get(ctx, key).Int64()
	log.Println("r", r, "err", err)
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return r, nil
}

func (rd *RedisStorage) Set(ctx context.Context, key string, expire int64) (int64, error) {
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
	count, err := incrBy.Run(ctx, rd.client, keys, values...).Int64()
	if err != nil {
		return 0, err
	}
	return count, nil
}
