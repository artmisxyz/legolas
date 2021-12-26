package database

import (
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"github.com/go-redis/redis/v8"
)

func NewRedis(conf configs.Config) redis.UniversalClient {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
	})
}
