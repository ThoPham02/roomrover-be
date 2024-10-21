package model

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ ContractRedisInterface = (*ContractRedis)(nil)

type ContractRedis struct {
	redisClient
}
type ContractRedisInterface interface {
	IncreaseContractCode(ctx context.Context, key string, ttl time.Duration) (count int, err error)
}

func NewContractRedisClient(redis *redis.Client) ContractRedis {
	return ContractRedis{
		*newCacheClient(redis),
	}
}

func (m *ContractRedis) IncreaseContractCode(ctx context.Context, key string, ttl time.Duration) (count int, err error) {
	data, err := m.redisConnect.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	m.redisConnect.ExpireNX(ctx, key, ttl)
	return int(data), nil
}
