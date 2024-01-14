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

type ChatRepo interface {
	CreateStream(ctx context.Context, c *Chat) error
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
