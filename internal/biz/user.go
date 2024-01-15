package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
}

type Message struct {
	Username string
	Topic    string
	Message  string
}

type Subscribe struct {
	Topic   string
	Message string
}

type UserRepo interface {
	//订阅主题
	Subscribe(ctx context.Context, s *Subscribe) error
	//发送消息
	SendMessage(ctx context.Context, m *Message) error
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

// 订阅主题
func (uc *UserUsecase) Subscribe(ctx context.Context, topic string) (*Subscribe, error) {
	s := &Subscribe{
		Topic: topic,
	}
	err := uc.ur.Subscribe(ctx, s)
	if err != nil {
		return nil, err
	}
	return &Subscribe{
		Topic:   s.Topic,
		Message: s.Message,
	}, nil
}

// 发送消息
func (uc *UserUsecase) SendMessage(ctx context.Context, username string, topic string, message string) (*Message, error) {
	m := &Message{
		Username: username,
		Topic:    topic,
		Message:  message,
	}
	err := uc.ur.SendMessage(ctx, m)
	if err != nil {
		return nil, err
	}
	return &Message{
		Username: m.Username,
		Topic:    m.Topic,
		Message:  m.Message,
	}, nil
}
