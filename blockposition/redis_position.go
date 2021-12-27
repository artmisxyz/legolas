package blockposition

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type redisPositionHolder struct {
	client      redis.UniversalClient
	startKey    string
	currentKey  string
	finishKey   string
	startBlock  uint64
	finishBlock uint64
}

func (r *redisPositionHolder) Update(u uint64) error {
	return r.client.Set(context.Background(), r.currentKey, u, 0).Err()
}

func (r *redisPositionHolder) Create() error {
	err := r.client.Set(context.Background(), r.startKey, r.startBlock, 0).Err()
	if err != nil {
		return err
	}
	err = r.client.Set(context.Background(), r.currentKey, r.startBlock, 0).Err()
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), r.finishKey, r.finishBlock, 0).Err()
}

func (r *redisPositionHolder) Exists() bool {
	return r.client.Exists(context.Background(), r.currentKey).Val() == 1
}

func (r *redisPositionHolder) Read(point Point) (uint64, error) {
	switch point {
	case Start:
		return r.client.Get(context.Background(), r.startKey).Uint64()
	case Current:
		return r.client.Get(context.Background(), r.currentKey).Uint64()
	case Finish:
		return r.client.Get(context.Background(), r.finishKey).Uint64()
	}
	return 0, fmt.Errorf("invlaid input")
}

func (r *redisPositionHolder) Incr() error {
	return r.client.Incr(context.Background(), r.currentKey).Err()
}

func NewRedisPositionHolder(inspector string, startBlock, finishBlock uint64, client redis.UniversalClient) BlockPositionHolder {
	return &redisPositionHolder{
		client:      client,
		startKey:    fmt.Sprintf("%s:block:start", inspector),
		currentKey:  fmt.Sprintf("%s:block:position", inspector),
		finishKey:   fmt.Sprintf("%s:block:finish", inspector),
		startBlock:  startBlock,
		finishBlock: finishBlock,
	}
}
