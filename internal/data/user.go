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

// Subscribe 订阅主题
func (r *userRepo) Subscribe(ctx context.Context, s *biz.Subscribe) error {
	r.data.rdb.Subscribe(ctx, s.Topic)
	return nil
}

// SendMessage 发送消息
func (r *userRepo) SendMessage(ctx context.Context, m *biz.Message) error {
	r.data.rdb.Publish(ctx, m.Topic, m.Message)
	return nil
}
