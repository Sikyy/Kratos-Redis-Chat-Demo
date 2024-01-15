package biz

import (
	"context"

	v1 "Redis-chat/api/chat/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Chat struct {
	StreamName string
	Key        string
	Value      string
}

type ConsumerGroup struct {
	GroupName string
	Stream    string
}

type Consumer struct {
	GroupName    string
	Stream       string
	ConsumerName string
}

type ChatRepo interface {
	CreateStream(ctx context.Context, c *Chat) error
	CreateConsumerGroup(ctx context.Context, c *ConsumerGroup) error
	AddConsumer(ctx context.Context, c *Consumer) error
	DelConsumer(ctx context.Context, c *Consumer) error
	DelConsumerGroup(ctx context.Context, c *ConsumerGroup) error
}

type ChatUsecase struct {
	repo ChatRepo
	log  *log.Helper
}

func NewChatUsecase(repo ChatRepo, logger log.Logger) *ChatUsecase {
	return &ChatUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ChatUsecase) CreateStream(ctx context.Context, streamname string, key string, value string) (*Chat, error) {
	c := &Chat{
		StreamName: streamname,
		Key:        key,
		Value:      value,
	}
	err := uc.repo.CreateStream(ctx, c)
	if err != nil {
		return nil, err
	}
	return &Chat{
		StreamName: c.StreamName,
		Key:        c.Key,
		Value:      c.Value,
	}, nil
}

func (uc *ChatUsecase) CreateConsumerGroup(ctx context.Context, stream string, groupname string) (*ConsumerGroup, error) {
	c := &ConsumerGroup{
		GroupName: groupname,
		Stream:    stream,
	}
	err := uc.repo.CreateConsumerGroup(ctx, c)
	if err != nil {
		return nil, err
	}
	return &ConsumerGroup{
		GroupName: c.GroupName,
		Stream:    c.Stream,
	}, nil
}

func (uc *ChatUsecase) DelConsumerGroup(ctx context.Context, stream string, groupname string) (*ConsumerGroup, error) {
	c := &ConsumerGroup{
		GroupName: groupname,
		Stream:    stream,
	}
	err := uc.repo.DelConsumerGroup(ctx, c)
	if err != nil {
		return nil, err
	}
	return &ConsumerGroup{
		GroupName: c.GroupName,
		Stream:    c.Stream,
	}, nil
}

func (uc *ChatUsecase) AddConsumer(ctx context.Context, stream string, groupname string, consumername string) (*Consumer, error) {
	c := &Consumer{
		GroupName:    groupname,
		Stream:       stream,
		ConsumerName: consumername,
	}
	err := uc.repo.AddConsumer(ctx, c)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		GroupName:    c.GroupName,
		Stream:       c.Stream,
		ConsumerName: c.ConsumerName,
	}, nil
}

func (uc *ChatUsecase) DelConsumer(ctx context.Context, stream string, groupname string, consumername string) (*Consumer, error) {
	c := &Consumer{
		GroupName:    groupname,
		Stream:       stream,
		ConsumerName: consumername,
	}
	err := uc.repo.DelConsumer(ctx, c)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		GroupName:    c.GroupName,
		Stream:       c.Stream,
		ConsumerName: c.ConsumerName,
	}, nil
}
