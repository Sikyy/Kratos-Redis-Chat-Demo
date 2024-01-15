package data

import (
	"Redis-chat/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type chatRepo struct {
	data *Data
	log  *log.Helper
}

func NewChatRepo(data *Data, logger log.Logger) biz.ChatRepo {
	return &chatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *chatRepo) CreateStream(ctx context.Context, c *biz.Chat) error {
	// 实现创建 Redis Stream 的逻辑
	r.data.rdb.XAdd(ctx, &redis.XAddArgs{
		// 用于指定 Stream 的名字，如果不存在则创建
		//表示创建一个名为 streamName 的 Stream
		//插入值为 message:hello 的消息
		Stream: c.StreamName,
		Values: map[string]interface{}{c.Key: c.Value},
	})
	return nil
}

func (r *chatRepo) CreateConsumerGroup(ctx context.Context, c *biz.ConsumerGroup) error {
	// 实现创建 Redis Consumer Group 的逻辑
	//传入Streamname 和 Groupname
	//创建一个名为 Groupname 的消费者组
	//设置消费者组的起始位置为 $，表示从最新的消息开始消费
	r.data.rdb.XGroupCreate(ctx, c.Stream, c.GroupName, "$")
	return nil
}

func (r *chatRepo) AddConsumer(ctx context.Context, c *biz.Consumer) error {
	//创建一个消费者并将其添加到指定的消费者组中
	r.data.rdb.XGroupCreateConsumer(ctx, c.Stream, c.GroupName, c.ConsumerName)
	return nil
}
