package data

import (
	"Redis-chat/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedis, NewChatRepo, NewUserRepo)

type Data struct {
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rdb: rdb}, cleanup, nil
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Username:     "", // no username set
		Password:     "", // no password set
		DB:           0,  // use default DB
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	fmt.Println("redis connect success")
	return rdb
}
