package ratelimit

import (
	"context"
)

type KeyStorage interface {
	TTL(ctx context.Context, key string) string
	Get(ctx context.Context, key string) (int64, error)
	Set(ctx context.Context, key string, expire int64) (int64, error)
}
