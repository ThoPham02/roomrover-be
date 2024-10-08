package model

import (
	"github.com/go-redis/redis/v8"
)

var _ redisInterface = (*redisClient)(nil)

type redisInterface interface {
	// Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	// Get(ctx context.Context, key string) (interface{}, error)
}

type redisClient struct {
	redisConnect *redis.Client
}

func newCacheClient(redis *redis.Client) *redisClient {
	return &redisClient{
		redisConnect: redis,
	}
}

// func (m *redisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
// 	if ttl < 0 {
// 		return ErrTTL
// 	} else {
// 		return m.redisConnect.Set(ctx, key, value, ttl).Err()
// 	}
// }

// func (m *redisClient) Get(ctx context.Context, key string) (interface{}, error) {
// 	if len(key) == 0 {
// 		return nil, ErrKey
// 	}
// 	return m.redisConnect.Get(ctx, key).Result()
// }
