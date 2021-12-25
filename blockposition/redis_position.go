package blockposition

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type redisPositionHolder struct {
	client     redis.UniversalClient
	key        string
	startBlock uint64
}

func (r *redisPositionHolder) Update(u uint64) error {
	return r.client.Set(context.Background(), r.key, u, 0).Err()
}

func (r *redisPositionHolder) Create() error {
	return r.client.Set(context.Background(), r.key, r.startBlock, 0).Err()
}

func (r *redisPositionHolder) Exists() bool {
	return r.client.Exists(context.Background(), r.key).Val() == 1
}

func (r *redisPositionHolder) Read() (uint64, error) {
	return r.client.Get(context.Background(), r.key).Uint64()
}

func (r *redisPositionHolder) Incr() error {
	return r.client.Incr(context.Background(), r.key).Err()
}

func NewRedisPositionHolder(inspector string, startBlock uint64, client redis.UniversalClient) BlockPositionHolder {
	return &redisPositionHolder{
		client:     client,
		key:        fmt.Sprintf("%s:block:position", inspector),
		startBlock: startBlock,
	}
}
