package data

import (
	"Redis-chat/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type userRepo struct {
	client *redis.Client
	data   *Data
	log    *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	client := data.rdb
	if client == nil {
		log.NewHelper(logger).Error("Redis client is nil")
		return nil
	}

	return &userRepo{
		client: client,
		data:   data,
		log:    log.NewHelper(logger),
	}
}

// 在 Redis 中创建消费者组
func (r *userRepo) CreateConsumerGroup(ctx context.Context, groupName string) error {
	// 实现创建 Redis 消费者组的逻辑
	if err := r.client.XGroupCreateMkStream(ctx, "yourStream", groupName, "$").Err(); err != nil {
		r.log.Error("Failed to create consumer group", "error", err)
		return err
	}
	return nil
}
