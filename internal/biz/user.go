package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type User struct {
	Setname  string //集合名称
	Username string //用户名
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

type Set struct {
	Setname      string
	SetPeopleNum int64
	SetPeople    []string
}

type UserRepo interface {
	//订阅主题
	Subscribe(ctx context.Context, s *Subscribe) error
	//发送消息
	SendMessage(ctx context.Context, m *Message) error
	//登录
	Login(ctx context.Context, u *User) error
	//登出
	Logout(ctx context.Context, u *User) error
	//获取集合内人数
	GetSetPeopleNum(ctx context.Context, s *Set) (*redis.IntCmd, error)
	//获取集合内人员
	GetSetPeople(ctx context.Context, s *Set) (*redis.StringSliceCmd, error)
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

// 登录
func (uc *UserUsecase) Login(ctx context.Context, setname string, username string) (*User, error) {
	u := &User{
		Setname:  setname,
		Username: username,
	}
	err := uc.ur.Login(ctx, u)
	if err != nil {
		return nil, err
	}
	return &User{
		Setname:  u.Setname,
		Username: u.Username,
	}, nil
}

// 登出
func (uc *UserUsecase) Logout(ctx context.Context, setname string, username string) (*User, error) {
	u := &User{
		Setname:  setname,
		Username: username,
	}
	err := uc.ur.Logout(ctx, u)
	if err != nil {
		return nil, err
	}
	return &User{
		Setname:  u.Setname,
		Username: u.Username,
	}, nil
}

// 获取集合内人数
func (uc *UserUsecase) GetSetPeopleNum(ctx context.Context, setname string) (*Set, error) {
	s := &Set{
		Setname: setname,
	}
	result, err := uc.ur.GetSetPeopleNum(ctx, s)
	if err != nil {
		return nil, err
	}
	return &Set{
		Setname:      s.Setname,
		SetPeopleNum: result.Val(),
	}, nil
}

// 获取集合内人员
func (uc *UserUsecase) GetSetPeople(ctx context.Context, setname string) (*Set, error) {
	s := &Set{
		Setname: setname,
	}
	result, err := uc.ur.GetSetPeople(ctx, s)
	if err != nil {
		return nil, err
	}
	return &Set{
		Setname:   s.Setname,
		SetPeople: result.Val(),
	}, nil
}
